import './App.css';
import { OpenDirectoryDialog, OpenFileDialog } from "../wailsjs/go/main/App";

function App() {
    return (
        <div id="App">
            <div id="input" className="input-box">
                <button className="btn" onClick={OpenDirectoryDialog}>OpenDirectoryDialog</button>
                <button className="btn" onClick={OpenFileDialog}>OpenFileDialog</button>
            </div>
        </div>
    )
}

export default App
