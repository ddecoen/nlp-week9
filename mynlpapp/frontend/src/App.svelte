<script>
  import logo from './assets/images/logo-universal.png'
  import {Greet} from '../wailsjs/go/main/App.js'



  let inputText = "";
  let outputText = "";

  const validTerms = [
    "break", "case", "chan", "const", "continue", "defer", "else", "fallthrough",
    "for", "func", "go", "goto", "if", "import", "interface", "map", "package",
    "range", "return", "select", "struct", "switch", "type", "var"
  ];

  const sillyTerms = [
    "blueberry", "dreams", "unicorn", "rainbow", "chocolate", "fairy", "moonbeam",
    "sparkle", "glitter", "bubble", "whimsy", "twinkle", "fizz", "flutter", "giggle"
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

<main>
  <div class="container">
    <div class="header-section">
        <!-- svelte-ignore a11y-img-redundant-alt -->
        <img src={heroImage} alt="Hero Image" class="hero-image" />
        <div>
            <h2>Week 9 Assignment: Access A Database</h2>
            <p>
                This assignment mimicks a desktop application that allows the user to ask questions of an LLM and receive a response.  The functionality of posing a question to an LLM and receiving an answer has been replaced with a database lookup.
            </p>
            <h2>How to Use</h2>
            <p>
                Enter text and click evaluate.  If the term exists in the database, the definition will be returned.  If the term does not exist in the database, an error message will be returned.  Note that the database search is case sensitive and requires an exact match.
            </p>
            <p>
              The "Valid Term" button on the right will randomly select a term that exists in the database.  Note this has been hard-coded.
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
