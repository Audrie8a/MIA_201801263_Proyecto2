import React, { useState } from 'react';
import Axios from 'axios';
import '../assets/css/Perfil.css';

const Perfil = (props) => {
    const [Username, setUser] = useState('')
    const [Password, setPassword] = useState('')
    const [Nombre, setNombre] = useState('')
    const [Apellido, setApellido] = useState('')
    const [FechaNac, setFechaNac] = useState('')
    const [Correo, setCorreo] = useState('')
    const [Foto, setFoto] = useState('')
    const Primero = "../assets/imagenes/";
    const [toggleState, setToggleState] = useState(0);
    const toggleTab = (index) => {
        setToggleState(index);
        if (index === 2) {

            setUser(props.usuario.Username);
            setPassword(props.usuario.Password);
            setNombre(props.usuario.Nombre);
            setApellido(props.usuario.Apellido);
            setFechaNac(props.usuario.FechaNac);
            setCorreo(props.usuario.Correo);
            setFoto(props.usuario.Foto);
        }

    }

    const handlerChange = (e) => {
        if (e.target.name === 'Password') { setPassword(e.target.value) }
        else if (e.target.name === 'Nombre') { setNombre(e.target.value) }
        else if (e.target.name === 'Apellido') { setApellido(e.target.value) }
        else if (e.target.name === 'FechaNac') { setFechaNac(e.target.value) }
        else if (e.target.name === 'Correo') { setCorreo(e.target.value) }
        else if (e.target.name === 'Foto') { setFoto(e.target.files[0].name) }
    }

    const submitEditar = () => {
        Axios.post('http://localhost:4000/updateUsuario',
            {
                Username: Username,
                Password: Password,
                Nombre: Nombre,
                Apellido: Apellido,
                FechaNac: FechaNac,
                Correo: Correo,
                Foto: Foto
            }).then((response) => {
                alert(response.data.Mensaje)
                
            })

    };
    return (
        <React.Fragment>
            
            <div>
                <img src={`data:image/gif, base64,${Primero+props.usuario.Foto}`} alt={props.usuario.Username}/>
                
                <h1>Bienvenido {props.usuario.Username}</h1>
                <br />
                
            </div>


            <div className="containerPerfil">
                <div className="bloc-tabs2">
                    <div className={toggleState === 1 ? "tabs2 active-tabs2" : "tabs2"}
                        onClick={() => toggleTab(1)}>Datos</div>
                    <div className={toggleState === 2 ? "tabs2 active-tabs2" : "tabs2"}
                        onClick={() => toggleTab(2)}>Editar</div>
                </div>
                <div className={toggleState === 1 ? "content active-content" : "content"}>

                    <div className="content active-content">
                        <div className="container3" style={{ marginTop: '20px' }}>
                            <div className="row">
                                <div className="col-lg-12">

                                    <table className="table table-bordered">
                                        <thead>
                                            <tr>
                                                <th>Username</th>
                                                <th>Password</th>
                                                <th>Nombre</th>
                                                <th>Apellido</th>
                                                <th >Tier</th>
                                                <th>Fecha Nacimiento</th>
                                                <th>Fecha Registro</th>
                                                <th>Correo</th>

                                            </tr>
                                        </thead>
                                        <tbody>
                                            <tr>
                                                <td >{props.usuario.Username}</td>
                                                <td>{props.usuario.Password}</td>
                                                <td>{props.usuario.Nombre}</td>
                                                <td>{props.usuario.Apellido}</td>
                                                <td >{props.membresia}</td>
                                                <td>{props.usuario.FechaNac}</td>
                                                <td>{props.usuario.FechaRegistro}</td>
                                                <td>{props.usuario.Correo}</td>
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
                                            <td><input type="text" name="Password" value={Password} onChange={handlerChange} /></td>
                                            <td><input type="text" name="Nombre" value={Nombre} onChange={handlerChange} required /></td>
                                            <td><input type="text" name="Apellido" value={Apellido} onChange={handlerChange} required /></td>
                                            <td><input type="date" name="FechaNac" min="1920-01-01" max="2003-05-07" onChange={handlerChange} required /></td>
                                            <td><input type="text" name="Correo" value={Correo} onChange={handlerChange} required /></td>
                                            <td><input type="file" name="Foto" onChange={handlerChange} required /></td>
                                        </tr>
                                    </tbody>
                                </table>
                                <br />
                                <button onClick={submitEditar}>Editar</button>
                            </div>
                        </div>
                       
                    </div>
                </div>

            </div>
        </React.Fragment>
    );


};

export default Perfil;