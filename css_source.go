package main

func stylesheetSource() string {
	return `
	/*  This file is GENERATED by the audiovisio tool
	 *
	 */ It will NOT be overwritten by the tool.

	body {
		font: 14px Arial, sans-serif;
		margin: 0px;
	}
	
	header {
		padding: 10px 20px;
		background: #589cb4;
		color: #3c444c;
	}
	
	header h1 {
		font-size: 24px;
	}
	
	.container {
		width: 100%;
		background: #157394;
	}
	
	nav,
	section {
		float: left;
		padding: 10px;
		min-height: 170px;
		box-sizing: border-box;
	}
	
	section {
		width: 80%;
	}
	
	nav {
		width: 20%;
		background: #bcb4ac;
	}
	
	nav ul {
		list-style: none;
		line-height: 24px;
		padding: 0px;
	}
	
	nav ul li a {
		color: #333;
	}
	
	.clearfix:after {
		content: ".";
		display: block;
		height: 0;
		clear: both;
		visibility: hidden;
	}
	
	footer {
		background: #3c444c;
		text-align: center;
		padding: 5px;
		color: white;
	}

`
}
