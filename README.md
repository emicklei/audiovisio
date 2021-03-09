# audiovisio - self-playing presentations

## intro

Audiovisio is a experimental tool to create HTML based presentations from full image slides for which each slide has its own audio file. 
Using a simple configuration file, you specify the title and order of each slide.
The audiovisio tool will generate all the webpage resources such that you can host the standalone presentation.

## usage

The example configuratiob below is taken from the test folder.
There are 2 slides, 4 images and 2 audio files.
The audio files are created using Google Text to Speech (see [gspeech](https://github.com/emicklei/gspeech)).

### generate

    audiovisio -c test.yaml -o test.html

```
title: Episode One
header: Episode One
footer: More like this 

# defaults
images-width: 950
images-height: 600

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

trailer:
  image: trailer.png
  ```