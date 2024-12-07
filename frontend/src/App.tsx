import {useState} from 'react';
import pp from './assets/images/pp.png';
import './App.css';
import {Cmd} from "../wailsjs/go/main/App";
import { EventsOn } from '../wailsjs/runtime/runtime';



function App() {
    const [resultText, setResultText] = useState("");
    const [countProcessedNodes, setCountProcessedNodes] = useState(0);
    const [countToProcess, setCountToProcess] = useState(0);

    EventsOn('app:directorySelected', (data) => {
        setResultText(data);
    });

    EventsOn('app:nodeToProcess', (data) => {
        setCountToProcess(data+1);
    });
    
    EventsOn('app:nodeProcessed', () => {
        setCountProcessedNodes(prevCount => prevCount + 1);
    });

    EventsOn('app:StartProcessing', () => {
        setResultText("Processing...");
    });

    return (
        <div id="App">
            <img src={pp} id="pp" alt="pp"/>
            <div id="countProcessedNodes" className="countProcessedNodes">Files Processed: {countProcessedNodes} / Files to Process:{countToProcess}</div>
            <div id="result" className="result">{resultText}</div>
            {resultText.includes("/") && <button onClick={() => Cmd(resultText)}>Open in Finder</button>}
        </div>
    )
}

export default App
