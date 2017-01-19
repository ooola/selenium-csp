#!/usr/bin/env python2.7

from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.desired_capabilities import DesiredCapabilities

desired_cap = {'browser': 'Chrome', 'browser_version': '55.0', 'os': 'Windows', 'os_version': '10', 'resolution': '1024x768'}

driver = webdriver.Remote(
    command_executor='http://--YourStuffHere--@hub.browserstack.com:80/wd/hub',
    desired_capabilities=desired_cap)

driver.get("http://YourEC2Here.compute.amazonaws.com:8080/")
driver.implicitly_wait(3)

print 'initial title: ' + driver.title
driver.find_element_by_id("title1").click()
driver.implicitly_wait(3)
print 'title after 1st click: ' + driver.title


driver.find_element_by_id("title2").click()
driver.implicitly_wait(3)
print 'title after 2nd click: ' + driver.title
print driver.get_log("browser")

#driver.save_screenshot("foo.png")

driver.quit()
