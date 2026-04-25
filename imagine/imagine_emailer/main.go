package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"golang.org/x/net/context"
)

const maxUploadSize = 10 << 20 // 10 MB

func main() {
	http.HandleFunc("/send_image", handleSendImage)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Photo booth service listening on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func handleSendImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "Request too large or malformed", http.StatusBadRequest)
		return
	}

	// --- Parse fields ---
	name := strings.TrimSpace(r.FormValue("name"))
	email := strings.TrimSpace(r.FormValue("email"))
	if name == "" || email == "" {
		http.Error(w, "Missing required fields: name, email", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Missing image field", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read image", http.StatusInternalServerError)
		return
	}

	// Detect content type from header or sniff from bytes
	contentType := header.Header.Get("Content-Type")
	if contentType == "" {
		contentType = http.DetectContentType(imageData)
	}
	ext := extensionFromContentType(contentType, header.Filename)

	// --- Send via SES ---
	fromAddr := os.Getenv("SES_FROM_ADDRESS")
	if fromAddr == "" {
		http.Error(w, "SES_FROM_ADDRESS not configured", http.StatusInternalServerError)
		return
	}

	if err := sendPhotoEmail(r.Context(), fromAddr, name, email, imageData, contentType, ext); err != nil {
		log.Printf("SES send error: %v", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	log.Printf("Photo sent to %s <%s>", name, email)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ok","message":"Photo sent to %s"}`, email)
}

// sendPhotoEmail builds a raw MIME email with the photo embedded inline and sends via SES.
func sendPhotoEmail(ctx context.Context, from, name, toEmail string, imageData []byte, contentType, ext string) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("load AWS config: %w", err)
	}
	client := ses.NewFromConfig(cfg)

	raw, err := buildRawEmail(from, name, toEmail, imageData, contentType, ext)
	if err != nil {
		return fmt.Errorf("build email: %w", err)
	}

	_, err = client.SendRawEmail(ctx, &ses.SendRawEmailInput{
		RawMessage:   &types.RawMessage{Data: raw},
		Source:       aws.String(from),
		Destinations: []string{toEmail},
	})
	return err
}

// buildRawEmail constructs a multipart/related MIME message so the photo renders
// inline inside the HTML body (referenced via cid:photo) rather than as an attachment.
func buildRawEmail(from, name, toEmail string, imageData []byte, contentType, ext string) ([]byte, error) {
	var buf bytes.Buffer
	boundary := "TalosPhotoBoothBound77"

	// --- Top-level headers ---
	fmt.Fprintf(&buf, "From: Talos RIT Photo Booth <%s>\r\n", from)
	fmt.Fprintf(&buf, "To: %s <%s>\r\n", name, toEmail)
	fmt.Fprintf(&buf, "Subject: [Talos RIT] Your Imagine RIT photo is here!\r\n")
	fmt.Fprintf(&buf, "MIME-Version: 1.0\r\n")
	// multipart/related ties the HTML and its inline image resource together
	fmt.Fprintf(&buf, "Content-Type: multipart/related; boundary=%q; type=\"text/html\"\r\n", boundary)
	fmt.Fprintf(&buf, "\r\n")

	mw := multipart.NewWriter(&buf)
	mw.SetBoundary(boundary)

	// --- Part 1: HTML body — references the image as cid:photo ---
	htmlHeader := make(map[string][]string)
	htmlHeader["Content-Type"] = []string{"text/html; charset=UTF-8"}
	htmlPart, err := mw.CreatePart(htmlHeader)
	if err != nil {
		return nil, err
	}
	if _, err := io.WriteString(htmlPart, buildHTML(name)); err != nil {
		return nil, err
	}

	// --- Part 2: inline image with Content-ID: <photo> ---
	imgHeader := make(map[string][]string)
	imgHeader["Content-Type"] = []string{fmt.Sprintf("%s; name=\"photo%s\"", contentType, ext)}
	imgHeader["Content-Transfer-Encoding"] = []string{"base64"}
	imgHeader["Content-Disposition"] = []string{"inline"}
	imgHeader["Content-ID"] = []string{"<photo>"}
	imgPart, err := mw.CreatePart(imgHeader)
	if err != nil {
		return nil, err
	}
	encoded := base64.StdEncoding.EncodeToString(imageData)
	// Wrap at 76 chars per RFC 2045
	for len(encoded) > 76 {
		if _, err := fmt.Fprintf(imgPart, "%s\r\n", encoded[:76]); err != nil {
			return nil, err
		}
		encoded = encoded[76:]
	}
	if len(encoded) > 0 {
		fmt.Fprintf(imgPart, "%s\r\n", encoded)
	}

	mw.Close()
	return buf.Bytes(), nil
}

// buildHTML returns the branded Talos RIT email HTML for the given recipient name.
// The photo is referenced as cid:photo — filled in by the inline MIME part.
func buildHTML(name string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1.0">
<title>Talos RIT — Your Photo</title>
</head>
<body style="margin:0;padding:0;background-color:#F6F1E9;">
<table width="100%%" cellpadding="0" cellspacing="0" style="background-color:#F6F1E9;padding:36px 20px;">
  <tr><td align="center">
    <table width="560" cellpadding="0" cellspacing="0" style="max-width:560px;width:100%%;">

      <!-- Top orange bar -->
      <tr><td style="background-color:#F76C00;height:5px;"></td></tr>

      <!-- Header -->
      <tr>
        <td style="background-color:#ffffff;padding:28px 36px 20px;border-left:1px solid #E8E2D9;border-right:1px solid #E8E2D9;">
          <table width="100%%" cellpadding="0" cellspacing="0"><tr>
            <td>
              <div style="font-family:Arial,Helvetica,sans-serif;font-size:10px;font-weight:700;letter-spacing:4px;text-transform:uppercase;color:#F76C00;margin-bottom:8px;">
                Talos RIT &mdash; Imagine RIT 2026
              </div>
              <div style="font-family:Arial,Helvetica,sans-serif;font-size:28px;font-weight:800;color:#1A1A1A;line-height:1.2;">
                Your photo is ready,<br>
                <span style="color:#F76C00;font-style:italic;">%s</span>
              </div>
            </td>
            <td align="right" valign="middle" style="padding-left:16px;white-space:nowrap;">
              <div style="border:2px solid #F76C00;border-radius:4px;padding:8px 14px;display:inline-block;text-align:center;">
                <span style="font-family:Arial,Helvetica,sans-serif;font-size:9px;font-weight:700;letter-spacing:3px;text-transform:uppercase;color:#F76C00;line-height:1.6;">SE&nbsp;Senior<br>Project</span>
              </div>
            </td>
          </tr></table>
        </td>
      </tr>

      <!-- Polaroid photo frame -->
      <tr>
        <td style="background-color:#ffffff;padding:4px 36px 28px;border-left:1px solid #E8E2D9;border-right:1px solid #E8E2D9;">
          <table width="100%%" cellpadding="0" cellspacing="0"><tr><td align="center">
            <!-- polaroid card with stacked shadow -->
            <table cellpadding="0" cellspacing="0"
              style="background:#ffffff;
                     border:1px solid #D9D2C8;
                     box-shadow:3px 5px 0 #D0C9BE, 6px 10px 0 #E4DDD5;
                     margin:12px auto;">
              <!-- wide white mat top + sides -->
              <tr>
                <td style="padding:12px 12px 0 12px;background:#ffffff;">
                  <img src="cid:photo" alt="Your Imagine RIT photo" width="460"
                    style="display:block;width:100%%;max-width:460px;height:auto;" />
                </td>
              </tr>
              <!-- caption strip -->
              <tr>
                <td style="padding:16px 12px 20px;text-align:center;background:#ffffff;border-top:1px solid #F0EBE3;">
                  <span style="font-family:'Comic Sans MS',cursive,sans-serif;font-size:14px;color:#999999;letter-spacing:0.5px;">
                    Imagine RIT 2026 &nbsp;
                  </span>
                </td>
              </tr>
            </table>
          </td></tr></table>
        </td>
      </tr>

      <!-- Body copy -->
      <tr>
        <td style="background-color:#ffffff;padding:4px 36px 28px;border-left:1px solid #E8E2D9;border-right:1px solid #E8E2D9;">
          <p style="font-family:Arial,Helvetica,sans-serif;font-size:14px;line-height:1.8;color:#555555;margin:0 0 20px;">
            The team at Talos RIT hopes you had a great time at our booth!
            Thanks for stopping by and checking out what we&rsquo;re building at Talos RIT!
          </p>
          <!-- GitHub CTA -->
          <table cellpadding="0" cellspacing="0">
            <tr>
              <td style="background-color:#FDF8F2;border-left:3px solid #F76C00;padding:14px 18px;">
                <div style="font-family:Arial,Helvetica,sans-serif;font-size:10px;font-weight:700;letter-spacing:3px;text-transform:uppercase;color:#F76C00;margin-bottom:5px;">
                  Curious about the project?
                </div>
                <a href="https://github.com/talos-rit"
                  style="font-family:Arial,Helvetica,sans-serif;font-size:13px;color:#1A1A1A;text-decoration:none;font-weight:600;">
                  github.com/talos-rit &rarr;
                </a>
              </td>
            </tr>
          </table>
        </td>
      </tr>

      <!-- Footer -->
      <tr>
        <td style="background-color:#F6F1E9;border:1px solid #E8E2D9;border-top:none;padding:14px 36px;">
          <table width="100%%" cellpadding="0" cellspacing="0"><tr>
            <td style="font-family:Arial,Helvetica,sans-serif;font-size:10px;color:#BBBBBB;letter-spacing:2px;text-transform:uppercase;">
              Talos RIT &mdash; Software Engineering Senior Project &mdash; RIT
            </td>
            <td align="right" style="font-size:16px;color:#DDDDDD;">&#128247;</td>
          </tr></table>
        </td>
      </tr>

      <!-- Bottom orange bar -->
      <tr><td style="background-color:#F76C00;height:3px;"></td></tr>

    </table>
  </td></tr>
</table>
</body>
</html>`, name)
}

// extensionFromContentType returns a file extension (e.g. ".jpg") for the image.
func extensionFromContentType(ct, originalFilename string) string {
	if originalFilename != "" {
		if e := filepath.Ext(originalFilename); e != "" {
			return e
		}
	}
	mediaType, _, _ := mime.ParseMediaType(ct)
	switch mediaType {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "image/webp":
		return ".webp"
	case "image/heic", "image/heif":
		return ".heic"
	default:
		return ".jpg"
	}
}
