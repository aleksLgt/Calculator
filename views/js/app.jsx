class App extends React.Component {
    render() {
        return (<Main />)
    }
}

class Main extends React.Component {
    render(){
        return(
            <section>
                <div className="container">
                    <div id="display"></div>
                    <div className="buttons">
                        <div className="button">C</div>
                        <div className="button">/</div>
                        <div className="button">*</div>
                        <div className="button">&larr;</div>
                        <div className="button">7</div>
                        <div className="button">8</div>
                        <div className="button">9</div>
                        <div className="button">-</div>
                        <div className="button">4</div>
                        <div className="button">5</div>
                        <div className="button">6</div>
                        <div className="button">+</div>
                        <div className="button">1</div>
                        <div className="button">2</div>
                        <div className="button">3</div>
                        <div className="button">.</div>
                        <div className="button">(</div>
                        <div className="button">0</div>
                        <div className="button">)</div>
                        <div id="equal" className="button">=</div>
                    </div>
                </div>
            </section>
        )
    }
}

ReactDOM.render(<App />, document.getElementById('app'));