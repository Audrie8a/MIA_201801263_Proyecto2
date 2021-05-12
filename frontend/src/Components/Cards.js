import React from 'react'
import PropTypes from 'prop-types'
const Colores = () => {

}
const Cards = ({name, Imagen, Color}) =>(
<div className="single-Deporte" style={{
    backgroundColor: Color
}}>
    <h2>{name}</h2>
    <img src={`data:image/gif,base64,${Imagen}`} alt={name}/>
    <button className="Botones">Editar</button>
   
</div>
)

Cards.propTypes={
    name: PropTypes.string
}
export default Cards;