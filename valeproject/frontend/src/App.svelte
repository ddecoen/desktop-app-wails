<script>
  import logo from './assets/images/logo-universal.png'
  //import {Greet} from '../wailsjs/go/main/App.js'

  //let resultText = "Please enter your name below 👇"
  //let name
  let userText = "";
  let analysisResult = "";

  async function analyzeText() {
        const response = await fetch("/api/analyze_text", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ text: userText }),
        });

        if (response.ok){

          const analysisResultText = await response.text();
          //const data = await response.json();
          console.log("API response:", analysisResultText); //Add for debugging
          analysisResult = analysisResultText;
        } else {
          console.error("Failed to analyze text.");
        }
    }

    async function onBlur(){
      await analyzeText();
    }

  //function greet() {
    //Greet(name).then(result => resultText = result)
  //}
</script>

<main>
  <h1>Text Analyzer</h1>
  <img alt="Wails logo" id="logo" src="{logo}">
  <div>
  <label for ="user-text">Enter text or markdown:</label>
  <textarea id="user-text" bind:value={userText}></textarea>
  <button on:click={analyzeText}>Analyze</button>
  <div>


  {#if analysisResult}
    <div class = "result" id="analysis-result">
        <h3>Analysis Result:</h3>
        <pre>{analysisResult}</pre>
    </div>
  {/if}
</main>

<style>

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
