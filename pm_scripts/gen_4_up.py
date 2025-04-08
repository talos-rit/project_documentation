import shutil
import os

FOURUP = "4up"
MD_EXT = ".md"
FOURUP_PATH = 'C:/Users/alexv/OneDrive/Desktop/RIT/GCCIS/SWEN-561-562/project_planning/pm/4-ups'
TEMPLATE_FILENAME = '4up_template.md'

def main():
    print("Generating 4-UP document...")
    year = input("Year: ")
    month = input("Month: ")
    day = input("Day: ")
    new_filename = "_".join([year, month, day, FOURUP]) + MD_EXT
    new_file_path = "/".join([FOURUP_PATH, new_filename])
    template_path = "/".join([FOURUP_PATH, TEMPLATE_FILENAME])

    if os.path.isfile(template_path):
        print("file " + template_path + " exists!")
    else:
        print("file " + template_path + "not found :(")
    shutil.copy2(template_path, new_file_path)
    print("Created new file: " + new_filename + "!")
        # try: 
        #     shutil.copy2(template_path, "//?/" + new_file_path)
        # except:
        # print("Failed to move the script "+os.path.basename(template_path)+" to "+ new_file_path)

if __name__ == "__main__":
    main()