import React from 'react'
import '../assets/css/CargaMasiva.css'
import Axios from 'axios'
const CargaMasiva = props => {    

    //const [path, setpath] = useState('');
    const submitCarga = () => {
        const yaml = require('js-yaml')
        const obj = yaml.load(ArchivoY)
        var dato= JSON.stringify(obj, null, 2)
        console.log(dato)
        Axios.post('http://localhost:4000/CargarDatos',{Info:
        dato}).then((response)=>{
            alert(response.data.Mensaje)
        })
    }
    var ArchivoY ="";
    var AbrirArchivo= function (evt){
        
        const fileY= evt.target.files[0];
        const reader= new FileReader();
        let lector = e=>{
            ArchivoY= e.target.result;
        }
        lector = lector.bind(evt);
        reader.onload= lector;
        reader.readAsText(fileY);
    }

    return (
        <React.Fragment>
            <div className="form5">
                <input type="file" id="Carga" name="Archivo" onChange={evt => AbrirArchivo(evt)} />
                <button id="Boton" onClick={submitCarga}>Cargar</button>
            </div>


        </React.Fragment>


    );
}

export default CargaMasiva;