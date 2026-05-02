Bale Web Automation (Go + chromedp)
A simple automation example written in Go using the chromedp library to interact with the Bale Web chat interface.

This project demonstrates how to control a Chromium-based browser (Edge/Chrome), navigate to a Bale chat page, insert a message into the chat input field, and send it automatically at a specified interval.

Features
Browser automation using chromedp
Works with Microsoft Edge or Google Chrome
Automatically sends a message to a Bale chat
Configurable message text and sending interval
Runs continuously in the same browser tab
How it works
The program launches a Chromium-based browser, opens a Bale chat URL, finds the message input field in the DOM, inserts text via JavaScript, and clicks the send button programmatically.

Requirements
Go 1.20+
Chromium-based browser (Edge, Chrome, etc.)
Bale Web account logged in on the opened browser
Note
This project is intended for learning and experimentation with browser automation using Go and chromedp.
