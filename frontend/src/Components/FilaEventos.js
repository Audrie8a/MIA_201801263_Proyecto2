import React from 'react'
import PropTypes from 'prop-types'


const FilaEventos = ({ IdEvento, IdJornada,Nombre,Fecha, NombreLocal, NombreVisitante }) => {
    
    

    return (
        <React.Fragment>  
            <tr>  
            <td>{IdEvento}</td>         
            <td>{IdJornada}</td> 
            <td>{Nombre}</td> 
            <td>{Fecha}</td> 
            <td>{NombreLocal}</td> 
            <td>{NombreVisitante}</td> 
            </tr>
            
            
        </React.Fragment>
    );
}

FilaEventos.propTypes = {
    name: PropTypes.string
}
export default FilaEventos;