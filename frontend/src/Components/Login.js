import React, {  useState, useEffect } from 'react'
import "../assets/css/Login.css"
import Axios from 'axios'

function Login (){
    const [Username, setUser] = useState('')
    const [Password, setPassword] = useState('')
    const submitIngreso=()=>{
        Axios.post('http://localhost:4000/Login',
        {
            Username:Username,
            Password: Password
        }).then(()=>{
            alert ('Ingresando...')
        });
    };

    return (
        <React.Fragment>
            <div className="form">
                <label>Usuario</label>
                <input type="text" name="Username" onChange={(e) =>
                    setUser(e.target.value)
                } />
                <label>Contraseña</label>
                <input type="password" name="Password" onChange={(e) =>
                    setPassword(e.target.value)
                } />
                <button onClick={submitIngreso}>Ingresar</button>
            </div>
        </React.Fragment>



    );
}

export default Login;