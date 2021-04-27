//import logo from './assets/imagenes/logo.svg';
import './assets/css/App.css';
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";
import Inicio from './Components/InicioTabs';
import paginaUsuarios from './Components/PaginaUsuarios';
function App() {
  return (
    
    <div className="App">
      <Router>
        <Switch>
          <Route exact path="/" component={Inicio}/>
          <Route exact path="/Usuario" component={paginaUsuarios}/>
        </Switch>
      </Router>
    </div>

  );
}

export default App;
