// Este componente de React renderiza un modal que solicita a usuarios con roles específicos completar su perfil con información adicional, enviando los datos a la API correspondiente para su verificación.

import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../../context/AuthContext';
import '../Styles/dashboard.css';
import { profileService } from '../../services/ApiServices';
import { FaIdCard, FaBuilding, FaUser, FaNotesMedical } from 'react-icons/fa'; 

const CompleteProfileModal = () => {
    const { user, completeProfile } = useAuth();
    const navigate = useNavigate();

    
    const [formData, setFormData] = useState({});
    const [error, setError] = useState(null);
    const [isLoading, setIsLoading] = useState(false);

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setIsLoading(true);
        setError(null);
        
        const payload = {
            ...formData,
            user_id: user.id, 
        };

        try {
            console.log(`Enviando perfil de ${user.role}:`, payload);
            
            switch (user.role) {
                case 'commerce':
                    await profileService.completeCommerceProfile(payload);
                    break;
                case 'council':
                    await profileService.completeCouncilProfile(payload);
                    break;
                case 'disabled':
                    await profileService.completeDisabledProfile(payload);
                    break;
                default:
                    break;
            }

            alert("¡Perfil completado con éxito!");
            completeProfile(); 
            navigate('/login');

        } catch (err) {
            const errorMessage = err.response?.data?.error || 'No se pudo guardar la información.';
            setError(errorMessage);
            console.error("Error al completar el perfil:", err.response);
        } finally {
            setIsLoading(false);
        }
    };
    
    const renderRoleForm = () => {
        switch (user.role) {
            case 'commerce':
                return (
                    <>
                        <div className="form-group">
                            <input type="text" id="rif" name="rif" className="form-input" placeholder=" " onChange={handleChange} required />
                            <label htmlFor="rif" className="form-label">RIF del Comercio</label>
                            <FaBuilding className="input-icon" />
                        </div>
                        <div className="form-group">
                            <input type="text" id="boss_name" name="boss_name" className="form-input" placeholder=" " onChange={handleChange} required />
                            <label htmlFor="boss_name" className="form-label">Nombre del Encargado</label>
                            <FaUser className="input-icon" />
                        </div>
                        <div className="form-group">
                            <input type="text" id="boss_document" name="boss_document" className="form-input" placeholder=" " onChange={handleChange} required />
                            <label htmlFor="boss_document" className="form-label">Cédula del Encargado</label>
                            <FaIdCard className="input-icon" />
                        </div>
                    </>
                );
            case 'council':
                return (
                    <>
                        <div className="form-group">
                            <input type="text" id="leader_name" name="leader_name" className="form-input" placeholder=" " onChange={handleChange} required />
                            <label htmlFor="leader_name" className="form-label">Nombre del Jefe de Calle/Líder</label>
                            <FaUser className="input-icon" />
                        </div>
                        <div className="form-group">
                            <input type="text" id="leader_document" name="leader_document" className="form-input" placeholder=" " onChange={handleChange} required />
                            <label htmlFor="leader_document" className="form-label">Cédula del Jefe de Calle/Líder</label>
                            <FaIdCard className="input-icon" />
                        </div>
                    </>
                );
            case 'disabled':
                return (
                    <>
                        <div className="form-group">
                            <input type="text" id="document" name="document" className="form-input" placeholder=" " onChange={handleChange} required />
                            <label htmlFor="document" className="form-label">Cédula de Identidad</label>
                            <FaIdCard className="input-icon" />
                        </div>
                        <div className="form-group">
                            <textarea id="disability" name="disability" className="form-input" rows="3" placeholder=" " onChange={handleChange} required></textarea>
                            <label htmlFor="disability" className="form-label">Descripción de la Discapacidad</label>
                            <FaNotesMedical className="input-icon" /> 
                        </div>
                    </>
                );
            default:
                return <p>Este tipo de usuario no requiere datos adicionales.</p>;
        }
    };

    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <h2>Completa tu Perfil</h2>
                <p>Necesitamos unos datos adicionales para verificar tu cuenta de <strong>{user.role}</strong>.</p>
                <form onSubmit={handleSubmit}>
                    {renderRoleForm()}
                    {error && <p className="error-message" style={{textAlign: 'center'}}>{error}</p>}
                    <button type="submit" className="btn-main-action" disabled={isLoading}>
                        {isLoading ? 'Guardando...' : 'Guardar y Continuar'}
                    </button>
                </form>
            </div>
        </div>
    );
};

export default CompleteProfileModal;