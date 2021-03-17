# audiovisio - build a self-playing presentation

## intro

Audiovisio is a experimental tool to create HTML based presentations from full image slides for which each slide has its own audio file.
Using a simple configuration file, you specify the title and order of each slide.
The audiovisio tool will generate all the webpage resources such that you can host the standalone presentation.

## usage

The example configuration below is taken from the test folder.
There are 2 slides, 4 images and 2 audio files.
The audio files are created using Google Text to Speech (see [gspeech](https://github.com/emicklei/gspeech)).

### generate

    audiovisio -c test.yaml -o test.html

```
title: Test - Episode One
header: Episode One
footer: More like this
button-start-title: Start 

# transition in milliseconds
# can be overridden per slide
pause-before-audio: 500
pause-after-audio: 500

leader:  
  image: leader.png
  next: 1

slides:
# id can be any non-empty string but must be unique in this document
- id: 1
  title: The Hitch Hikers Guide of the Galaxy
# image must be in a browser compatible format
  image: test1.png
  sound: test1.mp3
  next: 2  

- id: 2
  title: Test 2
  image: test2.png
  sound: test2.mp3
  # override transition in milliseconds
  pause-before-audio: 1000

trailer:
  image: trailer.png
```

### Contributions

Fixes, suggestions, documentation improvements are all welcome.
Fork this project and submit small Pull requests.
Discuss larger ones in the Issues list.
You can also sponsor Melrōse via [Github Sponsors](https://github.com/sponsors/emicklei).

Software is licensed under [MIT](LICENSE).
&copy; 2020 [ernestmicklei.com](http://ernestmicklei.com)