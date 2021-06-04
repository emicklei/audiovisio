package audiovisio

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
   <body>
        <script type="text/javascript" src="{{.Config.JSInclude}}"></script>
        <div class="wrapper">

            <header class="header">
                {{.Config.HeaderContent}}
            </header>

            <article class="main">
                <div class="box">
                    <img id="image_holder" src="{{.Config.Leader.Image}}" />
                </div>
            </article>            
            <aside class="aside aside-1">
            
                <ol>
                    {{ range .Slides -}}
                    <li><a id="start_{{ .ID }}" href="#" onclick="play('{{.ID}}')">{{ .Title -}}</a></li> 
                    {{ end -}}
                </ol>                
                
            </aside>
            <aside class="aside aside-2"></aside>
            <footer class="footer">
            
                <p>{{.Config.FooterContent}}</p>

                <button id="watch" type="button" onclick="play('{{.FirstSlideID}}')">{{.Config.ButtonStartTitle}}</button>

                <audio controls id="audio_holder" onended="play_next()" hidden="true">
                    <source id="audio_source" src="" type="audio/mpeg"> Your browser does not support the audio tag.
                </audio>

            </footer>
        </div>

        <script type="text/javascript">
          function slides() {
              let slides = new Map();
              {{ range .Slides -}}
              slides.set('{{.ID}}', { 
                  "title": "{{.Title}}", 
                  "audio": "{{.Sound}}", 
                  "visio": "{{.Image}}" , 
                  "next": '{{.NextID}}' ,
                  "pause_before": {{.PauseBeforeAudio}} ,
                  "pause_after": {{.PauseAfterAudio}} });
              {{ end -}}
              slides.set('config.leader',{
                  "visio": "{{.Config.Leader.Image}}"
              });
              slides.set('config.trailer',{
                  "visio": "{{.Config.Trailer.Image}}"
              });
              return slides;
          }
        </script>

    </body>
</html>
`
}
