 Source Generated with Decompyle++
# File: EduInfomationScan v1.1.0.pyc (Python 3.9)

Unsupported opcode: WITH_EXCEPT_START
import sys
import os
import threading
from urllib.parse import urlparse
from urllib.parse import urljoin
from bs4 import BeautifulSoup
import re
import collections
import urllib.parse as urllib
import time
import random
from multiprocessing.dummy import Pool
import queue
from selenium import webdriver
from selenium.webdriver.common.desired_capabilities import DesiredCapabilities
import ddddocr
from PIL import Image
from PIL import ImageFile
ImageFile.LOAD_TRUNCATED_IMAGES = True
from selenium.webdriver.common.by import By

def getExterLinks(bs, exterurl):
Warning: block stack is not empty!
    for link in bs.find_all('a', re.compile('^(http|www)((?!' + urlparse(exterurl).netloc + ').)*$'), **('href',)):
        link.attrs['href'] = urllib.parse.quote(link.attrs['href'], '?=&:/@', **('safe',))
        if link.attrs['href'] is not None and link.attrs['href'] not in externalLinks:
            queueExt.put(link.attrs['href'])
            externalLinks.append(link.attrs['href'])
            print(link.attrs['href'])
            continue
            return None


def getInterLinks(bs, current_url, thread_driver):
Unsupported opcode: RERAISE
    site_domain = urlparse(current_url).netloc
    a_list = thread_driver.find_elements_by_tag_name('a')
# WARNING: Decompyle incomplete


def download_file(bs, current_url, thread_driver):
Unsupported opcode: RERAISE
    current_url_domain = '{}'.format(urlparse(current_url).netloc)
    file_elementlist = []
    filetpye = [
        '.doc',
        '.docx',
        '.xls',
        '.xlxs',
        '.pdf',
        '.txt',
        '.csv',
        'zip',
        '.7z',
        '.rar']
    cu_filename_list = []
    all_a_list = thread_driver.find_elements_by_tag_name('a')
# WARNING: Decompyle incomplete


def deepLinks(thread_driver):
Segmentation fault