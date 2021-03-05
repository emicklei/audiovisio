package main

func htmlTemplateSource() string {
	return `<!DOCTYPE html>
<!-- This file is GENERATED with the audiovisio tool
     It will be overwritten each time you run the tool with your configuration.
-->
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>{{.Config.Title}}</title>
    <link rel="stylesheet" href="{{.Config.CSSInclude}}" />
</head>
<!-- http://css3.bradshawenterprises.com/cfimg/ -->
<body>
    <script type="text/javascript" src="{{.Config.JSInclude}}"></script>
    <div class="container">
        <header>
            <h1>{{.Config.HeaderContent}}</h1>
        </header>
        <div class="wrapper clearfix">
            <nav>
                <button id="watch" type="button" onclick="play('{{.FirstSlideID}}')">Start presentation</button>
                <ul>
                    {{ range .Slides -}}
                    <li><a id="start_{{ .ID }}" href="#" onclick="play('{{.ID}}')">{{ .Title -}}</a></li> 
                    {{ end -}}
                </ul>
            </nav>
            <section>
                <img id="image_holder" src="leader.png" width="{{.Config.ImagesWidth}}px" height="{{.Config.ImagesHeight}}px" />
            </section>
        </div>
        <footer>
            <p>{{.Config.FooterContent}}</p>
            <audio controls id="audio_holder" onended="play_next()" hidden="true">
                <source id="audio_source" src="" type="audio/mpeg"> Your browser does not support the audio tag.
            </audio>            
        </footer>
    </div>

    <script type="text/javascript">
        function slides() {
            let slides = new Map();
            {{ range .Slides -}}
            slides.set('{{.ID}}', { "title": "{{.Title}}", "audio": "{{.Sound}}", "visio": "{{.Image}}" , "next": '{{.NextID}}' });
            {{ end -}}
            return slides;
        }
    </script>

</body>
</html>
`
}
