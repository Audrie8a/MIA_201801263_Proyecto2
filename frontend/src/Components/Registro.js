import React, {  useState} from 'react'
import "../assets/css/Registro.css"
import Axios from 'axios'

function Registro (){
    const [Username, setUser] = useState('')
    const [Password, setPassword] = useState('')
    const [Nombre, setNombre] = useState('')
    const [Apellido, setApellido] = useState('')
    const [FechaNac, setFechaNac] = useState('')
    const [Correo, setCorreo] = useState('')
    const [Foto, setFoto] = useState('')
    const submitRegistro=()=>{
        Axios.post('http://localhost:4000/Registro',
        {
            Username:Username,
            Password: Password,
            Nombre: Nombre,
            Apellido: Apellido,
            FechaNac: FechaNac,
            Correo: Correo,
            Foto: Foto
        }).then((response)=>{
            alert(response.data.Mensaje)
        })
        
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
                <label>Nombre</label>
                <input type="text" name="Nombre" onChange={(e) =>
                    setNombre(e.target.value)
                } />
                <label>Apellido</label>
                <input type="text" name="Apellido" onChange={(e) =>
                    setApellido(e.target.value)
                } />
                <label>Fecha Nacimiento</label>
                <input type="date" name="FechaNac" value="2018-07-22" min="1920-01-01" max="2003-05-07" onChange={(e) =>
                    setFechaNac(e.target.value)
                } />
                <label>Correo</label>
                <input type="text" name="Correo" onChange={(e) =>
                    setCorreo(e.target.value)
                } />
                <label>Foto</label>
                <input type="file" name="Foto" onChange={(e) =>
                    setFoto(e.target.value)
                } />
                <button onClick={submitRegistro}>Registrar</button>
            </div>
        </React.Fragment>



    );
}

export default Registro;