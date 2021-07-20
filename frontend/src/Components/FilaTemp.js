import React from 'react'
import PropTypes from 'prop-types'


const FilaTemp = ({ Username,Temp, Total }) => {
    
    

    return (
        <React.Fragment>  
            <tr>  
            <td>{Username}</td>         
            <td>{Temp}</td> 
            <td>{Total}</td> 
            </tr>
            
            
        </React.Fragment>
    );
}

FilaTemp.propTypes = {
    name: PropTypes.string
}
export default FilaTemp;