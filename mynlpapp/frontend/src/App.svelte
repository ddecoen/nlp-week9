<script>
 // import logo from './assets/images/logo-universal.png'
 // import {Greet} from '../wailsjs/go/main/App.js'
  import {CheckTermEmbed} from '../wailsjs/go/main/App.js'
// @ts-ignore
  import helperImage from './assets/images/helperImage.png'


  let inputText = "";
  let outputText = "";

  const validTerms = [
    "break", "case", "chan", "const", "continue", "defer", "else", "fallthrough",
    "for", "func", "go", "goto", "if", "import", "interface", "map", "package",
    "range", "return", "select", "struct", "switch", "type", "var"
  ];

  const sillyTerms = [
    "blackberry", "tofu", "buffalo", "storm", "vanilla", "elves", "sunshine",
    "shine", "pizza", "smiley", "jokester", "supernova", "soda", "gusty", "laugh"
  ];

  function evaluateText() {
    console.log("evaluateText function called with input:", inputText);
    CheckTermEmbed(inputText).then(result => {
        console.log("Received result from backend:", result);
        outputText = result;
    });
}

  function loadSample(isValid) {
    if (isValid) {
      inputText = validTerms[Math.floor(Math.random() * validTerms.length)];
    } else {
      inputText = sillyTerms[Math.floor(Math.random() * sillyTerms.length)];
    }
  }


</script>

<style>
  .container {
      padding: 20px;
      display: flex;
      flex-direction: column;
      align-items: center;
  }

  .header-section {
    display: flex;
    align-items: center; 
    gap: 20px; 
}


  .helper-image {
    max-width: 200px; 
    margin-right: 20px;
  }

  .input-section {
      display: flex;
      align-items: start;
      width: 90%;
  }

  textarea {
      flex: 1;
      height: 200px;
      margin-right: 20px;
  }

  .sample-buttons {
      display: flex;
      flex-direction: column;
  }

  button {
      margin-bottom: 20px;
  }

  .output {
      width: 90%;
      border: 1px solid #ccc;
      padding: 10px;
      white-space: pre-line;
      text-align: left; 
  }

  .rules {
      text-align: left;
      font-size: 0.8rem;
  }
</style>


<div class="container">
  <div class="header-section">
      <!-- svelte-ignore a11y-img-redundant-alt -->
       <img src={helperImage} alt="Helper Image" class="helper-image" />
       <div>
           <h2>Week 9 Assignment: Access A Database for Natual Language Processing</h2>
           <p>
               This assignment mimicks a desktop application that allows the user to ask questions of an LLM and receive a response.  The functionality of posing a question to an LLM and receiving an answer has been replaced with a database lookup.
            </p>
        </div>
    </div>
    <div class="input-section">
        <textarea bind:value={inputText} placeholder="Enter your text here..."></textarea>
        <div class="sample-buttons">
            <button on:click={() => loadSample(true)}>Valid Term</button>
            <button on:click={() => loadSample(false)}>Invalid Term</button>
        </div>
    </div>
    <div class="button-container">
        <button on:click={evaluateText}>Evaluate</button>
    </div>
    <div class="output">{outputText}</div>
    <div class="rules">
        <h3>Valid Terms</h3>
        <p>
          "break", "case", "chan", "const", "continue", "defer", "else", "fallthrough",
          "for", "func", "go", "goto", "if", "import", "interface", "map", "package",
          "range", "return", "select", "struct", "switch", "type", "var"
      </p>
  </div>
  </div>
