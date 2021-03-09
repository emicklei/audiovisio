
	var current = 1;

	function play(id) {
		current = id
		let slide = slides().get(id)
		if (slide === undefined) {
			console.log("no such slide id",id)
			return
		}
	
		var watch = document.getElementById("watch");
		watch.hidden = true;
	
		var img = document.getElementById("image_holder");
		img.src = slide.visio;
	
		var aud = document.getElementById("audio_holder");            
		aud.hidden = false;   
		
		var source = document.getElementById("audio_source");
		source.src = slide.audio;
		aud.load();
	
		var playPromise = aud.play();
		// In browsers that don’t yet support this functionality,
		// playPromise won’t be defined.
		if (playPromise !== undefined) {
			playPromise.then(function() {
			}).catch(function(error) {
			// Re-playback failed.
			console.log("failed to re-play",error);
			});
		}  
	}
		
	function play_next() {
		let slide = slides().get(current)
		if (slide === undefined) {
			console.log("no such slide id",current)
			return
		}
		let next = slide.next
		if (next === undefined || next.length === 0) {
			end_presentation();
			return;
		}    
		play(next);
	}         
	
	function end_presentation() {
		var img = document.getElementById("image_holder");
		img.src = "trailer.png";
	
		var watch = document.getElementById("watch");
		watch.hidden = false;
	
		var aud = document.getElementById("audio_holder");
		// do not empty the src field
		aud.hidden = true;
	}