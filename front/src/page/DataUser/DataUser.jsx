// Este componente de React obtiene y muestra el perfil del usuario autenticado, incluyendo detalles específicos basados en su rol.

import React, { useState, useEffect } from 'react';
import { useAuth } from '../../context/AuthContext';
import { userService } from '../../services/ApiServices';
import '../Styles/dashboard.css';
import {  FaUserCircle, FaInfoCircle, FaMapMarkerAlt, FaFileAlt } from 'react-icons/fa';

const MyData = () => {
    const { user } = useAuth();
    const [profile, setProfile] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);

useEffect(() => {
    const fetchProfile = async () => {
        if (!user?.id) {
            setIsLoading(false);
            setError("Usuario no autenticado.");
            return;
        }
        try {
            const response = await userService.getProfileById(user.id);

            const mergedProfile = {
                ...response.data,
                role: user.role,           
                role_data: user.role_data,
                profileCompleted: user.profileCompleted 
            };

            setProfile(mergedProfile);
        } catch (err) {
            console.error("Error al cargar el perfil:", err);
            setError("No se pudo cargar la información del perfil.");
        } finally {
            setIsLoading(false);
        }
    };
    fetchProfile();
}, [user]);


    const capitalizeRole = (roleString) => {
        if (!roleString) return 'N/A';
        return roleString.charAt(0).toUpperCase() + roleString.slice(1);
    };

    const renderRoleSpecificDetails = () => {
        if (!profile?.role_data) return null;

        const roleData = profile.role_data;

        switch (profile.role) {
            case 'commerce':
                return (
                    <>
                        <h3><FaFileAlt /> Detalles de Comercio</h3>
                        <p><strong>RIF:</strong> {roleData.rif || 'N/A'}</p>
                        <p><strong>Nombre del Encargado:</strong> {roleData.boss_name || 'N/A'}</p>
                        <p><strong>Cédula del Encargado:</strong> {roleData.boss_document || 'N/A'}</p>
                    </>
                );
            case 'council':
                return (
                    <>
                        <h3><FaFileAlt /> Detalles de Consejo Comunal</h3>
                        <p><strong>Nombre del Líder:</strong> {roleData.leader_name || 'N/A'}</p>
                        <p><strong>Cédula del Líder:</strong> {roleData.leader_document || 'N/A'}</p>
                    </>
                );
            case 'disabled':
                return (
                    <>
                        <h3><FaFileAlt /> Detalles de Persona con Discapacidad</h3>
                        <p><strong>Documento:</strong> {roleData.document || 'N/A'}</p>
                        <p><strong>Discapacidad:</strong> {roleData.disability || 'N/A'}</p>
                    </>
                );
            default:
                return null;
        }
    };

    if (isLoading) {
        return <div className="loading-spinner">Cargando tu perfil...</div>;
    }

    if (error || !profile) {
        return (
            <div className="dashboard-content">
                <header className="dashboard-header"><h1>Mi Perfil</h1></header>
                <p className="error-message">{error || "No se pudo cargar la información del perfil."}</p>
            </div>
        );
    }

    return (
        <div className="dashboard-content">
            <header className="dashboard-header">
                <h1>Mi Perfil</h1>
                <p>Consulta tus datos de usuario y detalles de tu rol.</p>
            </header>

            <div className="dashboard-card my-profile-card">
                <FaUserCircle className="profile-icon" />
                <h2>{profile.name || 'N/A'}</h2>
                <p><strong>Email:</strong> {profile.email || 'N/A'}</p>

                <div className="profile-basic-info">
                    <p><strong>Rol:</strong> {capitalizeRole(profile.role)}</p>
                    {profile.municipality?.name ? (
                        <p><strong><FaMapMarkerAlt /> Municipio:</strong> {profile.municipality.name}</p>
                    ) : (
                        <p><strong><FaMapMarkerAlt /> Municipio:</strong> N/A</p>
                    )}
                </div>

                {profile.role !== 'user' && !profile.profileCompleted && (

                    <div className="profile-incomplete-warning">
                        <FaInfoCircle />
                        <p>Tu perfil aún no está completo. Por favor, completa los datos adicionales en el modal.</p>
                    </div>
                )}

                <div className="role-details-section">
                    {renderRoleSpecificDetails()}
                </div>
            </div>
        </div>
    );
};

export default MyData;