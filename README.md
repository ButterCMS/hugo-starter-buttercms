# Hugo + ButterCMS Starter Project

Live demo: [https://hugo-starter-buttercms.netlify.app/](https://hugo-starter-buttercms.netlify.app/) 

This Hugo starter project fully integrates with dynamic sample content from your ButterCMS account, including main menu, pages, blog posts, categories, and tags, and all with a beautiful, custom theme. All of the included sample content is automatically created in your account dashboard when you 
[sign up for a free trial](https://buttercms.com/join/) of ButterCMS.

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/ButterCMS/hugo-starter-buttercms#BUTTERCMS_API_TOKEN=Your_ButterCMS_Token_Here&BUTTERCMS_PREVIEW=false)


## 1) Installation

First, create a virtual environment and install dependencies by running the 
below commands.

```bash
$ git clone https://github.com/ButterCMS/hugo-starter-buttercms.git
$ cd hugo-starter-buttercms
$ go mod download 
```

### 2) Set ButterCMS API Token

To fetch your ButterCMS content, add your API token as an environment variable. 

```bash
$ echo 'BUTTERCMS_API_TOKEN=your_token' >> .env
```

### 3) Fetch Data from ButterCMS

To fetch your ButterCMS content run:

```bash
$ go run github.com/ButterCMS/hugo-starter-buttercms/prebuild
```

### 5) Run Hugo Local Development Server

To view the app in the browser, you'll need to run the local development server:

```bash
$ hugo server
```

Congratulations! Your starter project is now live at: [http://localhost:1313/](http://localhost:1313/)

### 6) Deploy on Netlify

Our starter app can be deployed to Netlify with the click of a button:

1. Create a Heroku account at https://app.netlify.com.
2. Click the button below and fill in an app name. Then click "Deploy to Netlify".

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/ButterCMS/hugo-starter-buttercms#BUTTERCMS_API_TOKEN=Your_ButterCMS_Token_Here&BUTTERCMS_PREVIEW=false)

### 7) Webhooks

In order for your deployed app to pull in content changes from your ButterCMS account, you'll need to follow your hosting provider's steps for setting up webhooks. The ButterCMS webhook settings are located at https://buttercms.com/webhooks/. 

### 8) Previewing Draft Changes

By default, your starter project is set up to allow previewing of draft changes saved in your ButterCMS.com account. To disable this functionality, set the following value in your **.env** file: `BUTTERCMS_PREVIEW=false`.
