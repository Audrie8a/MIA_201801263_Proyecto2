import React, {  useState, useEffect } from 'react'
import "../assets/css/Login.css"
import Axios from 'axios'
import {useHistory} from 'react-router-dom';

const Login = props =>{
    var history = useHistory();
    const [Username, setUser] = useState('')
    const [Password, setPassword] = useState('')
    const submitIngreso=()=>{
        Axios.post('http://localhost:4000/Login',
        {
            Username:Username,
            Password: Password
        }).then((response)=>{
            alert (response.data.Mensaje)
            if (response.data.Mensaje==="Acceso Concedido!"){
                if (Username==="201801263"){
                    history.push("/Admin/"+Username );  
                }else{
                    history.push("/Usuario/"+Username );  
                }             
            }
        })
        
    };

    return (
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


    );
}

export default Login;