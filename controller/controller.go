package controller

import (
	"fmt"

	"github.com/aksanart/go-graphDB-neo4j-crud/config"
	"github.com/aksanart/go-graphDB-neo4j-crud/service"
	"github.com/gin-gonic/gin"
)

func GetPerson(sc service.ControllerInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.Block{
			Try: func() {
				result, err := sc.FindAll()
				if err != nil {
					config.Throw(err.Error())
				}
				c.JSON(200, gin.H{
					"Code":        "S1000",
					"Status":      "SUCCESS",
					"Description": "Anjay sukses",
					"Value":       result,
				})
			},
			Catch: func(e config.Exception) {
				c.AbortWithStatusJSON(400, gin.H{
					"Code":        "E9999",
					"Status":      "FAILED",
					"Description": e,
				})
			},
			Finally: func() {
			},
		}.Do()
	}
}

func AddPerson(sc service.ControllerInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.Block{
			Try: func() {
				if c.PostForm("name") == "" || c.PostForm("age") == "" {
					config.Throw("user and following_name is required")
				}
				result, err := sc.AddPerson(c.PostForm("name"), c.PostForm("age"))
				if err != nil {
					config.Throw(err.Error())
				}
				c.JSON(200, gin.H{
					"Code":        "S1000",
					"Status":      "SUCCESS",
					"Description": "Anjay sukses",
					"Value":       result,
				})
			},
			Catch: func(e config.Exception) {
				c.AbortWithStatusJSON(400, gin.H{
					"Code":        "E9999",
					"Status":      "FAILED",
					"Description": e,
				})
			},
			Finally: func() {
			},
		}.Do()
	}
}

func Follow(sc service.ControllerInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.Block{
			Try: func() {
				if c.PostForm("user") == "" || c.PostForm("following_name") == "" {
					config.Throw("user and following_name is required")
				}
				result, err := sc.Follow(c.PostForm("user"), c.PostForm("following_name"))
				if err != nil {
					config.Throw(err.Error())
				}
				c.JSON(200, gin.H{
					"Code":        "S1000",
					"Status":      "SUCCESS",
					"Description": "Anjay sukses",
					"Value":       result,
				})
			},
			Catch: func(e config.Exception) {
				c.AbortWithStatusJSON(400, gin.H{
					"Code":        "E9999",
					"Status":      "FAILED",
					"Description": e,
				})
			},
			Finally: func() {
			},
		}.Do()
	}
}

func UnFollow(sc service.ControllerInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.Block{
			Try: func() {
				if c.PostForm("user") == "" || c.PostForm("following_name") == "" {
					config.Throw("user and following_name is required")
				}
				result, err := sc.UnFollow(c.PostForm("user"), c.PostForm("following_name"))
				if err != nil {
					config.Throw(err.Error())
				}
				c.JSON(200, gin.H{
					"Code":        "S1000",
					"Status":      "SUCCESS",
					"Description": "Anjay sukses",
					"Value":       result,
				})
			},
			Catch: func(e config.Exception) {
				c.AbortWithStatusJSON(400, gin.H{
					"Code":        "E9999",
					"Status":      "FAILED",
					"Description": e,
				})
			},
			Finally: func() {
			},
		}.Do()
	}
}

func ViewGraph(c *gin.Context) {
	host := "neo4j://localhost:7687"
	user := "public"
	pass := "public"
	fmt.Fprintf(c.Writer, `
	<!doctype html>
	<html>
    <head>
    <script src="https://cdn.neo4jlabs.com/neovis.js/v1.5.0/neovis.js"></script>
        <title>Graph DB view</title>
        <style type="text/css">
            html, body {
                font: 16pt arial;
            }
            #viz {
                width: 900px;
                height: 700px;
                border: 1px solid lightgray;
                font: 12pt arial;
            }
        </style>
    </head>
    <body onload="draw()">
        <div id="viz"></div>
    </body>
	</html>
	<script type="text/javascript">
        var viz;
        function draw() {
            var config = {
                container_id: "viz",
                server_url: "`+host+`",
                server_user: "`+user+`",
                server_password: "`+pass+`",
                labels: {
                    "Person": {
                        "caption": "name",
                        "size": "pagerank",
                        "community": "community",
                        "title_properties": [
                            "name",
                            "age"
                        ]
                    }
                },
                relationships: {
                    "Follow": {
                        "thickness": "weight",
                        "caption": true,
						"color":'Green'
                    }
                },
				arrows: true,
                initial_cypher: "USE graphDB MATCH (n) OPTIONAL MATCH (n)-[r]-(m) RETURN n, r,m"
            };
            viz = new NeoVis.default(config);
            viz.render();
        }
    </script>
		`)
}
