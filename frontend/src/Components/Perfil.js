import React from 'react';
import '../assets/css/Perfil.css'

const Perfil = (props) => {
    return (
        <React.Fragment>
            <div>
                <h1>Bienvenido {props.usuario.Username}</h1>
            </div>
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
                                    <th>Tier</th>
                                    <th>Fecha Nacimiento</th>
                                    <th>Fecha Registro</th>
                                    <th>Correo</th>

                                </tr>
                            </thead>
                            <tbody>
                                <tr>
                                    <td>{props.usuario.Username}</td>
                                    <td>{props.usuario.Password}</td>
                                    <td>{props.usuario.Nombre}</td>
                                    <td>{props.usuario.Apellido}</td>
                                    <td>{props.membresia}</td>
                                    <td>{props.usuario.FechaNac}</td>
                                    <td>{props.usuario.FechaRegistro}</td>
                                    <td>{props.usuario.Correo}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </React.Fragment>
    );


};

export default Perfil;