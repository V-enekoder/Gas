// Este componente de React permite a los usuarios crear un nuevo reporte sobre una entrega, seleccionando la entrega afectada, el tipo de problema y añadiendo una descripción detallada.

import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import '../Styles/dashboard.css';
import { useAuth } from '../../context/AuthContext';
import { reportService, deliveryService } from '../../services/ApiServices';


const CreateReport = () => {
    const navigate = useNavigate();
    const { user } = useAuth();

    const [userDeliveries, setUserDeliveries] = useState([]);
    const [reportTypes, setReportTypes] = useState([]);
    
    const [deliveryId, setDeliveryId] = useState('');
    const [typeId, setTypeId] = useState('');
    const [description, setDescription] = useState('');

    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            if (!user?.id) { setIsLoading(false); return; }

            try {
                const [deliveriesRes, typesRes] = await Promise.all([
                    deliveryService.getByUserId(user.id),
                    reportService.getReportTypes()
                ]);

                const deliveriesData = deliveriesRes.data || [];
                const typesData = typesRes.data || [];

                setUserDeliveries(deliveriesData);
                setReportTypes(typesData);

                if (deliveriesData.length > 0) setDeliveryId(deliveriesData[0].id);
                if (typesData.length > 0) setTypeId(typesData[0].id);

            } catch (err) {
                console.error("Error al cargar datos para el formulario:", err);
                setError("No se pudieron cargar los datos necesarios. Inténtalo de nuevo.");
            } finally {
                setIsLoading(false);
            }
        };
        fetchData();
    }, [user]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!deliveryId || !typeId || !description) {
            alert('Por favor, completa todos los campos.');
            return;
        }
        setIsLoading(true);
        setError(null);
        
        const reportPayload = {
            user_id: user.id, 
            delivery_id: parseInt(deliveryId, 10),
            type_id: parseInt(typeId, 10),
            description: description,
            report_state_id: 1,
        };

        try {
            await reportService.create(reportPayload); 
            alert("¡Reporte creado con éxito!");
            navigate('/reports');
        } catch (err) {
            setError(err.response?.data?.error || "No se pudo crear el reporte.");
        } finally {
            setIsLoading(false);
        }
    };

    if (isLoading) {
        return <div className="loading-spinner">Cargando datos del formulario...</div>
    }

    return (
        <div className="dashboard-content">
            <header className="dashboard-header">
                <h1>Crear Nuevo Reporte</h1>
                <p>Describe el problema que tuviste con una de tus entregas.</p>
            </header>
            <div className="dashboard-card form-card">
                <form onSubmit={handleSubmit}>
                    <div className="form-group">
                        <label htmlFor="delivery">Selecciona la Entrega Afectada</label>
                        <select id="delivery" value={deliveryId} onChange={(e) => setDeliveryId(e.target.value)} required className="form-select" disabled={userDeliveries.length === 0}>
                            {userDeliveries.length > 0 ? (
                                userDeliveries.map(d => (
                                    <option key={d.id} value={d.id}>
                                        Entrega DEL-{d.id} (de la Orden ORD-{d.order.id})
                                    </option>
                                ))
                            ) : (
                                <option disabled>No tienes entregas para reportar</option>
                            )}
                        </select>
                    </div>
                    <div className="form-group">
                        <label htmlFor="reportType">Tipo de Reporte</label>
                        <select id="reportType" value={typeId} onChange={(e) => setTypeId(e.target.value)} required className="form-select" disabled={reportTypes.length === 0}>
                            {reportTypes.map(rt => (
                                <option key={rt.id} value={rt.id}>{rt.name}</option>
                            ))}
                        </select>
                    </div>
                    <div className="form-group">
                        <label htmlFor="description">Descripción del Problema</label>
                        <textarea id="description" rows="6" placeholder="Por favor, sé lo más detallado posible..." value={description} onChange={(e) => setDescription(e.target.value)} required></textarea>
                    </div>
                    {error && <p className="error-message">{error}</p>}
                    <button type="submit" className="btn-main-action" disabled={isLoading || userDeliveries.length === 0}>
                        {isLoading ? 'Enviando...' : 'Enviar Reporte'}
                    </button>
                </form>
            </div>
        </div>
    );
};

export default CreateReport;