# Hugo + ButterCMS Starter Project

Live demo: [https://hugo-starter-buttercms.netlify.app/](https://hugo-starter-buttercms.netlify.app/) 

This Hugo starter project fully integrates with dynamic sample content from your ButterCMS account, including main menu, pages, blog posts, categories, and tags, and all with a beautiful, custom theme. All of the included sample content is automatically created in your account dashboard when you 
[sign up for a free trial](https://buttercms.com/join/) of ButterCMS.

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/ButterCMS/hugo-starter-buttercms#BUTTERCMS_API_TOKEN=Your_ButterCMS_Token_Here&BUTTERCMS_PREVIEW=false)


## Installation

First, install the required dependencies:

### Installing Go (1.22 or higher)

**Ubuntu/Debian:**
```bash
$ sudo apt update
$ sudo apt install golang-go

**macOS (using Homebrew):**
```bash
$ brew install go
```

### Installing Hugo (0.120.4 or higher)

**Ubuntu/Debian:**
```bash
$ sudo apt update
$ sudo apt install hugo
```

**macOS (using Homebrew):**
```bash
$ brew install hugo
```

### Setting up the project

Clone the repository and download dependencies:

```bash
$ git clone https://github.com/ButterCMS/hugo-starter-buttercms.git
$ cd hugo-starter-buttercms
$ go mod download 
```

### Set ButterCMS API Token

To fetch your ButterCMS content, add your API token as an environment variable. 

```bash
$ echo 'BUTTERCMS_API_TOKEN=your_token' >> .env
```

### Fetch Data from ButterCMS

To fetch your ButterCMS content run:

```bash
$ go run github.com/ButterCMS/hugo-starter-buttercms/prebuild
```

*Because Hugo is a static site generator, you'll need to run this command for if you make changes to your Butter content and want to pull it into your app. For this reason, we recommend tapping into our webhook architecture, which will allow you to set up your deployed project to automatically show draft changes.*

### Run Hugo Local Development Server

To view the app in the browser, you'll need to run the local development server:

```bash
$ hugo server
```

Congratulations! Your starter project is now live at: [http://localhost:1313/](http://localhost:1313/)

### Deploy on Netlify

Our starter app can be deployed to Netlify with the click of a button:

1. Create a Heroku account at https://app.netlify.com.
2. Click the button below and fill in an app name. Then click "Deploy to Netlify".

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/ButterCMS/hugo-starter-buttercms#BUTTERCMS_API_TOKEN=Your_ButterCMS_Token_Here&BUTTERCMS_PREVIEW=false)

### Webhooks

In order for your deployed app to pull in content changes from your ButterCMS account, you'll need to follow your hosting provider's steps for setting up webhooks. The ButterCMS webhook settings are located at https://buttercms.com/webhooks/. 

### Previewing Draft Changes

By default, your starter project is set up to allow previewing of draft changes saved in your ButterCMS.com account. To disable this functionality, set the following value in your **.env** file: `BUTTERCMS_PREVIEW=false`.
