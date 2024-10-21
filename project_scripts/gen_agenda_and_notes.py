import shutil
import os

AGENDA = "agenda"
NOTES = "notes"
MD_EXT = ".md"
MEETING_NOTES_PATH = 'C:/Users/alexv/OneDrive/Desktop/RIT/GCCIS/SWEN-561-562/project_planning/meeting_notes'
AGENDA_TEMPLATE_FILENAME = 'agenda_template.md'
NOTES_TEMPLATE_FILENAME = 'notes_template.md'

def main():
    print("Generating meeting notes documents...")
    year = input("Year: ")
    month = input("Month: ")
    day = input("Day: ")
    agenda_filename = "_".join([year, month, day, AGENDA]) + MD_EXT
    notes_filename = "_".join([year, month, day, NOTES]) + MD_EXT
    agenda_file_path = "/".join([MEETING_NOTES_PATH, agenda_filename])
    agenda_template_path = "/".join([MEETING_NOTES_PATH, AGENDA_TEMPLATE_FILENAME])
    
    notes_file_path = "/".join([MEETING_NOTES_PATH, notes_filename])
    notes_template_path = "/".join([MEETING_NOTES_PATH, NOTES_TEMPLATE_FILENAME])
    

    if os.path.isfile(agenda_template_path):
        print("file " + agenda_template_path + " exists!")
    else:
        print("file " + agenda_template_path + "not found :(")
    shutil.copy2(agenda_template_path, agenda_file_path)
    print("Created new file: " + agenda_file_path + "!")

    if os.path.isfile(notes_template_path):
        print("file " + notes_template_path + " exists!")
    else:
        print("file " + notes_template_path + "not found :(")
    shutil.copy2(notes_template_path, notes_file_path)
    print("Created new file: " + notes_filename + "!")

if __name__ == "__main__":
    main()