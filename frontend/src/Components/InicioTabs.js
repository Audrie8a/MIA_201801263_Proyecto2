import React, {useState} from "react";
import "../assets/css/InicioTabs.css";
import Login from "./Login";
import Registro from './Registro';

function InicioTabs(){
    const [toggleState, setToggleState]=useState(1);
    const toggleTab=(index)=>{
        setToggleState(index);
    }
    return(
        <React.Fragment>
            <div className="Titulo" >
                <h1> </h1>
            </div>
            <div className="container">
                <div className="bloc-tabs">
                    <div className={toggleState=== 1 ? "tabs active-tabs" : "tabs"}
                    onClick={()=>toggleTab(1)}>Login</div>
                    <div className={toggleState=== 2 ? "tabs active-tabs" : "tabs"} 
                    onClick={()=>toggleTab(2)}>Registro</div>
                </div>
                <div className={toggleState=== 1 ? "content active-content" : "content"}>

                    <div className="content active-content">
                       
                        <Login/>
                    </div>
                </div>
                <div className={toggleState=== 2 ? "content active-content" : "content"}>

                    <div className="content active-content">
                       
                        <Registro/>
                    </div>
                </div>
            </div>

            
                        
        </React.Fragment>
    );
}

export default InicioTabs;