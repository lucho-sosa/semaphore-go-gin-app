package main

func initializeRoutes(){
  //handle the index route
  router.GET("/", showIndexPage)
  //handle GET request for /article/view/some_article_id
  router.GET("/article/view/:article_id", getArticle)
}
