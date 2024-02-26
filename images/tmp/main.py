import os
import random
import string
import shutil
def copy_files(directory, total_files):
    jpg_files = [filename for filename in os.listdir(directory) if filename.endswith('.jpg')]
    existing_files_count = len(jpg_files)

    for i in range(1, total_files + 1):
        random_file = random.choice(jpg_files)
        original_path = os.path.join(directory, random_file)
        new_name = str(i) + '.jpg'
        new_path = os.path.join(directory, new_name)

        shutil.copyfile(original_path, new_path)


directory = '.'  # 当前目录
total_files = 432
copy_files(directory, total_files)
