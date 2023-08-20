<script>
 // import logo from './assets/images/logo-universal.png'
 // import {Greet} from '../wailsjs/go/main/App.js'
  import {CheckTermEmbed} from '../wailsjs/go/main/App.js'
// @ts-ignore
  import helperImage from './assets/images/helperImage.png'


  let inputText = "";
  let outputText = "";

  let validTermsVisible = false; // Initialize visibility state

  const validTermList = [
    "break", "case", "chan", "const", "continue", "defer", "else", "fallthrough",
    "for", "func", "go", "goto", "if", "import", "interface", "map", "package",
    "range", "return", "select", "struct", "switch", "type", "var"
  ];

  let terms = ["break", "case", "chan", "const", "continue", "defer", "else", "fallthrough",
    "for", "func", "go", "goto", "if", "import", "interface", "map", "package",
    "range", "return", "select", "struct", "switch", "type", "var"];

  let validTerms = [];

  function evaluateText() {
    console.log("evaluateText function called with input:", inputText);
    CheckTermEmbed(inputText).then(result => {
        console.log("Received result from backend:", result);
        outputText = result;
    });
}

  function handleButtonClick(){
    if(validTermsVisible){
      validTerms =[]; // clear the valid terms list
    } else{
      validTerms = terms.filter(term => ValidTermEmbed(term));
    }
    validTermsVisible = !validTermsVisible //toggle visibility state
    
  }

  function ValidTermEmbed(term){
    return validTermList.includes(term);
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
    height: 200px; /* Adjust the height as needed */
    font-size: 20px; /* Increase the font size for a bigger text area */
  }

  button {
      margin-bottom: 20px;
  }

  .sample-buttons {
    display: flex;
    flex-direction: column;
  }

  .styled-button {
    padding: 10px 20px;
    margin-bottom: 10px;
    background-color: #007bff;
    color: #ffffff;
    border: none;
    cursor: pointer;
    border-radius: 5px;
    font-size: 10px;
  }

  .valid-terms-list {
    text-align: center;
    margin-bottom: auto; /* Pushes the list to the bottom */
    padding-bottom: 20px; /* Optional: Add some padding */
  }

  .vertical-column {
    column-count: 3; /* Adjust the number of columns as needed */
    column-gap: 20px; /* Adjust the gap between columns */
    text-align: left; /* Align items to the left within columns */
  }

  .vertical-list li {
    margin-bottom: 5px;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    margin-bottom: 5px;
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
               Instructions - please enter a term in single quotes. The app is case-sensitive as well.
            </p>
        </div>
    </div>

    <div class="input-section">
      <textarea bind:value={inputText} placeholder="Enter your text here..."></textarea>
      <div class = "sample-buttons">
        <button on:click={handleButtonClick} class ="styled-button">Show Valid Terms</button>
      </div>
  </div>

    <div class="output">{outputText}</div>

    <div class="button-container">
        <button on:click={evaluateText}>Evaluate</button>
   
  {#if validTermsVisible}
    <div class="valid-terms-list">
      <h2>Valid Terms:</h2>
      <div class="vertical-column">
        <ul class="vertical-list">
          {#each validTerms as term}
            <li>{term}</li>
          {/each}
        </ul>
    </div>
   </div>
  {/if}
  <div class="rules">
  </div>
</div>
</div>

