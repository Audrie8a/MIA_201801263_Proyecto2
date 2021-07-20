import React from 'react'
import PropTypes from 'prop-types'


const FilaQuinelaUsuario = ({ Username, IdQuinela, Puntaje, Nombre, Local, Visitante, NombreVisitante,NombreLocal, ResultadoVisitante, ResultadoLocal, Fecha }) => {
    
    

    return (
        <React.Fragment>  
            <tr>  
            <td>{Username}</td>         
            <td>{IdQuinela}</td> 
            <td>{Puntaje}</td> 
            <td>{Nombre}</td> 
            <td>{NombreLocal}</td> 
            <td>{ResultadoLocal}</td> 
            <td>{Local}</td> 
            <td>{NombreVisitante}</td> 
            <td>{ResultadoVisitante}</td> 
            <td>{Visitante}</td> 
            <td>{Fecha}</td> 
            </tr>
            
            
        </React.Fragment>
    );
}

FilaQuinelaUsuario.propTypes = {
    name: PropTypes.string
}
export default FilaQuinelaUsuario;