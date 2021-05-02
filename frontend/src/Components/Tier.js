import React, {useState}from 'react'
import Select from 'react-select'
import '../assets/css/Tier.css'

const options=[
    {value: 1, label:'Gold'},
    {value: 2, label:'Silver'},
    {value: 3, label:'Bronze'},
]
const Tier = props =>{
    const [toggleState, setToggleState] = useState(0);
    const toggleTab = (index) => {
        setToggleState(index);   
    }
    return (
           <React.Fragment>
               
            <div>
                <h1>Hola {props.usuario.Username}</h1>
                <br />
            </div>


            <div className="containerTier">
                <div className="bloc-tabs3">
                    <div className={toggleState === 1 ? "tabs3 active-tabs3" : "tabs3"}
                        onClick={() => toggleTab(1)}>Asignar</div>
                    <div className={toggleState === 2 ? "tabs3 active-tabs3" : "tabs3"}
                        onClick={() => toggleTab(2)}>Editar</div>
                </div>
                <div className={toggleState === 1 ? "content active-content" : "content"}>

                    <div className="content active-content">
                        <div className="container4" style={{ marginTop: '20px' }}>
                            <div className="row">
                                <div className="col-lg-12">

                                    <table className="table table-bordered">
                                        <thead>
                                            <tr>
                                                <th>Membresia</th>

                                            </tr>
                                        </thead>
                                        <tbody>
                                            <tr>
                                                <td ><Select options={options}/></td>
                                                <td><button>Elegir</button></td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                            <div className={toggleState === 2 ? "content active-content" : "content"}>

                                <div className="content active-content">

                                </div>
                            </div>

                        </div>


                    </div>
                </div>
                <div className={toggleState === 2 ? "content active-content" : "content"}>

                    <div className="content active-content">
                        <div className="ContenedorEdit">
                            <div className="formEditado">
                                <table className="table table-bordered">
                                    <thead>
                                        <tr>
                                            <th>Contraseña</th>
                                            <th>Nombre</th>
                                            <th>Apellido</th>
                                            <th>Fecha Nacimiento</th>
                                            <th>Correo</th>
                                            <th>Foto</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
                                            
                                        </tr>
                                    </tbody>
                                </table>
                                <br />
                                <button >Editar</button>
                            </div>
                        </div>
                       
                    </div>
                </div>

            </div>
           </React.Fragment>


    );
}

export default Tier;