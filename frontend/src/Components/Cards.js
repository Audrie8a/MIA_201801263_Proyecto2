import React, { useState } from 'react'
import PropTypes from 'prop-types'
import Select from 'react-select'
import Axios from 'axios';
const options=[
    //"red", "pink", "aqua", "blue", "brown", "olive", "green", "teal", "yellow", "fuchsia", "lime"
    {value: 'red',label:'red'},
    {value: 'pink',label:'pink'},
    {value: 'aqua',label:'aqua'},
    {value: 'blue',label:'blue'},
    {value: 'brown',label:'brown'},
    {value: 'olive',label:'olive'},
    {value: 'green',label:'green'},
    {value: 'teal',label:'teal'},
    {value: 'yellow',label:'yellow'},
    {value: 'fuchsia', label:'fuchsia'},
    {value: 'lime', label:'lime'},
] 
const Cards = ({ name, Imagen, Color, Id }) => {
    
    const [Colors, setColor] = useState(Color)
    const [Foto, setFoto] = useState(Imagen)
    const submitRegistro=()=>{
        //if (valor===1){setColor('red');
        //}else if (valor===2){setColor('pink');
        //}else if (valor==3){setColor('aqua')
        //}else if (valor==4){setColor('blue')
        //}else if (valor==5){setColor('brown')
        //}else if (valor==6){setColor('olive')
        //}else if (valor==7){setColor('green')
        //}else if (valor==8){setColor('teal')
        //}else if (valor==9){setColor('yellow')
        //}else if (valor==10){setColor('fuchsia')
        //}else if (valor==11){setColor('lime')}
        
        
        
        Axios.post('http://localhost:4000/updateDeporte',
            {
                idDeporte: Id.toString(),
                Nombre: name,
                Imagen: Foto,
                Color: Colors
            }).then((response) => {
                alert(response.data.Mensaje)
                
            })


    }
    const [toggleState, setToggleState] = useState(1);
    const toggleTab = (index) => {
        setToggleState(index);
    }


    return (
        <React.Fragment>

            <div className="Titulo" >
                <h1> </h1>
            </div>
            <div className="containerCards">
                <div className="bloc-tabs">
                    <div className={toggleState === 1 ? "tabs active-tabs4" : "tabs"}
                        onClick={() => toggleTab(1)}>Deporte</div>
                    <div className={toggleState === 2 ? "tabs active-tabs4" : "tabs"}
                        onClick={() => toggleTab(2)}>Editar</div>
                    
                </div>
                <div className={toggleState === 1 ? "content active-content" : "content"}>

                    <div className="content active-content">
                        <div className="single-Deporte" style={{
                            backgroundColor: Color
                        }}>
                            <h2>{name}{Id}</h2>
                            <img src={`data:image/gif,base64,${Imagen}`} alt={name} />
                            

                        </div>

                    </div>
                </div>
                <div className={toggleState === 2 ? "content active-content" : "content"}>

                    <div className="content active-content">
                    <div className="form">
                <label>Color</label>
                <div clasName="SelectColor"><Select  options={options} onChange={(e) =>
                    setColor(e.value)
                } />
                </div>
                
                <label>Foto</label>
                <input type="file" name="Foto" onChange={(e) =>
                    setFoto(e.target.files[0].name)
                } />
                <button onClick={submitRegistro}>Editar</button>
            </div>
                    </div>
                </div>
               
            </div>


        </React.Fragment>
    );
}

Cards.propTypes = {
    name: PropTypes.string
}
export default Cards;