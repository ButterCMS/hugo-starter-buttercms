# Hugo + ButterCMS Starter Project

Live demo: [TBD: https://hugo-starter-buttercms-demo.netlify.com/](https://hugo-starter-buttercms-demo.netlify.com/) 

This Hugo starter project fully integrates with dynamic sample content from your ButterCMS account, including main menu, pages, blog posts, categories, and tags, and all with a beautiful, custom theme with already-implemented search functionality. All of the included sample content is automatically created in your account dashboard when you 
[sign up for a free trial](https://buttercms.com/join/) of ButterCMS.

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/ButterCMS/hugo-starter-buttercms)


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

### 6) Deploy on Heroku

Our starter app can be deployed to Netlify with the click of a button:

1. Create a Heroku account at https://app.netlify.com.
2. Click the button below and fill in an app name. Then click "Deploy to Netlify".

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/ButterCMS/hugo-starter-buttercms)

### 7.) Previewing and Draft Changes

Your starter app allows you to preview draft changes made to content in your
ButterCMS.com account by default; to see draft changes, add `PREVIEW_DRAFTS=true` env variable:

```bash
$ echo 'PREVIEW_DRAFTS=true' >> .env
```

You have to repeat step 3 to fetch data again.
