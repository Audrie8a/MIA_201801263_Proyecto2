import React, { useState } from "react";
import "../assets/css/PaginaUsuarios.css";
import Axios from 'axios';
import Perfil from './Perfil';
import Tier from './Tier'
function PaginaUsuarios({ match: { params: { id } } }) {
        const [FotoUrl, setFotoUrl]= useState('');
        const [Respuesta, setRespuesta] = useState(JSON);
        const [Membresia, setMembresia] = useState('');
        const [toggleState, setToggleState] = useState(1);
        
        const toggleTab = (index) => {
                setToggleState(index);
                if (index === 1 || index===2) {
                        Axios.post('http://localhost:4000/datosUsuario',
                                {
                                        Username: id
                                }).then((response) => {
                                        if(response.data.Tier===1){
                                                setMembresia("Gold");
                                        }else if(response.data.Tier===2){
                                                setMembresia("Silver");
                                        }else if(response.data.Tier===3){
                                                setMembresia("Bronze");
                                        }else {
                                                setMembresia("No registrado");
                                        }

                                        if(response.data.Foto===''){
                                                setFotoUrl("../assets/imagenes/NoImagen.jpg")
                                        }else{
                                                setFotoUrl(response.data.Foto)
                                        }

                                        setRespuesta(response.data);
                                })
                }

        }

        return (
                <React.Fragment>
                        <div className="Titulo" >
                                <h1>TODO DEPORTE GT</h1>
                        </div>
                        <div className="containerUsuario">
                                <div className="bloc-tabs">
                                        <div className={toggleState === 1 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(1)}>Perfil</div>
                                        <div className={toggleState === 2 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(2)}>Membresia</div>
                                        <div className={toggleState === 3 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(3)}>Quinela</div>
                                        <div className={toggleState === 4 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(4)}>Eventos</div>
                                        <div className={toggleState === 5 ? "tabs active-tabs" : "tabs"}
                                                onClick={() => toggleTab(5)}>Recompensas</div>
                                </div>
                                <div className={toggleState === 1 ? "content active-content" : "content"}>

                                        <div className="content active-content">

                                                <Perfil usuario={Respuesta} membresia={Membresia} FotoPath={FotoUrl} />

                                        </div>
                                </div>
                                <div className={toggleState === 2 ? "content active-content" : "content"}>

                                        <div className="content active-content">
                                                <Tier usuario={Respuesta}/>         
                                       
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
                        </div>



                </React.Fragment>

        );
};

export default PaginaUsuarios;