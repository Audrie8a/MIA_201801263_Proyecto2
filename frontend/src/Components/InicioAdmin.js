import React, { useState } from "react";
import "../assets/css/PaginaUsuarios.css";
import CargaMasiva from './CargaMasiva';

function InicioAdmin({ match: { params: { id } } }) {
        const [toggleState, setToggleState] = useState(1);
        
        const toggleTab = (index) => {
                setToggleState(index);               
        }

        return (
                <React.Fragment>
                        <div className="Titulo" >
                                <h1>TODO DEPORTE GT</h1>
                        </div>
                        <div className="containerUsuario">
                                <div className="bloc-tabs">
                                        <div className={toggleState === 1 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(1)}>Quinelas</div>
                                        <div className={toggleState === 2 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(2)}>Carga Masiva</div>
                                        <div className={toggleState === 3 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(3)}>Jornadas</div>
                                        <div className={toggleState === 4 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(4)}>Temporada</div>
                                        <div className={toggleState === 5 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(5)}>Recompensas</div>
                                        <div className={toggleState === 6 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(6)}>Deportes</div>
                                        <div className={toggleState === 7 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(7)}>Reportes</div>
                                </div>
                                <div className={toggleState === 1 ? "content active-content" : "content"}>

                                        <div className="content active-content">

                                                

                                        </div>
                                </div>
                                <div className={toggleState === 2 ? "content active-content" : "content"}>

                                        <div className="content active-content">
                                                <CargaMasiva usuario={id}/>            

                                        </div>
                                </div>
                                <div className={toggleState === 3 ? "content active-content" : "content"}>

                                        <div className="content active-content">


                                        </div>
                                </div>
                                <div className={toggleState === 4 ? "content active-content" : "content"}>

                                        <div className="content active-content">


                                        </div>
                                </div>
                                <div className={toggleState === 5 ? "content active-content" : "content"}>

                                        <div className="content active-content">


                                        </div>
                                </div>
                                <div className={toggleState === 6 ? "content active-content" : "content"}>

                                        <div className="content active-content">


                                        </div>
                                </div>
                                <div className={toggleState === 7 ? "content active-content" : "content"}>

                                        <div className="content active-content">


                                        </div>
                                </div>
                        </div>



                </React.Fragment>

        );
};

export default InicioAdmin;