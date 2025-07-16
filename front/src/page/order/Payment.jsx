// Este componente de React gestiona el proceso de confirmación de pago para una orden específica, registrando el pago y creando automáticamente una entrega asociada.

import { useParams, useNavigate } from 'react-router-dom';
import '../Styles/dashboard.css';
import React, { useState, useEffect } from 'react';
import { useAuth } from '../../context/AuthContext';
import { orderService, paymentService, deliveryService  } from '../../services/ApiServices'
import { FaCheckCircle } from 'react-icons/fa';

const Payment = () => {
    const { orderId } = useParams();
    const navigate = useNavigate();
    const { user } = useAuth();

    const [order, setOrder] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchOrderDetails = async () => {
            if (!orderId) { setIsLoading(false); return; }
            try {
                const response = await orderService.getById(orderId);
                setOrder(response.data);
            } catch (err) {
                setError("No se pudo encontrar la orden especificada.");
            } finally {
                setIsLoading(false);
            }
        };
        fetchOrderDetails();
    }, [orderId]);

    const handleConfirmPaymentAndCreateDelivery = async () => {
        if (!order) return;

        setIsLoading(true);
        setError(null);

        try {
            const paymentPayload = {
                user_id: user.id,
                order_id: parseInt(orderId, 10),
                quantity: order.total_price,
            };
            const paymentResponse = await paymentService.create(paymentPayload);
            const createdPayment = paymentResponse.data;

            const deliveryPayload = {
                order_id: parseInt(orderId, 10),
                payment_id: createdPayment.id,
            };
            await deliveryService.create(deliveryPayload);
            
            alert("¡Pago registrado y entrega programada con éxito!");
            navigate('/deliveries');

        } catch (err) {
            const errorMessage = err.response?.data?.error || "Ocurrió un error en el proceso.";
            setError(errorMessage);
        } finally {
            setIsLoading(false);
        }
    };
    
    if (isLoading && !order) {
        return <div className="loading-spinner">Cargando detalles del pedido...</div>;
    }

    if (!order) {
        return (
            <div className="dashboard-content">
                <header className="dashboard-header"><h1>Error</h1></header>
                <p className="error-message">{error || "No se pudo cargar la orden."}</p>
            </div>
        );
    }
    
    return (
        <div className="dashboard-content">
            <header className="dashboard-header">
                <h1>Confirmar Pago del Pedido</h1>
                <p>ID del Pedido: <span className="order-id-highlight">{orderId}</span></p>
            </header>
            
            <div className="dashboard-card confirmation-card">
                <h3>Resumen del Pago</h3>
                <div className="confirmation-details">
                    <div>
                        <span>Monto Total a Pagar:</span>
                        <strong>${order.total_price.toFixed(2)}</strong>
                    </div>
                </div>
                {error && <p className="error-message">{error}</p>}
                
                <button 
                    onClick={handleConfirmPaymentAndCreateDelivery} 
                    className="btn-main-action" 
                    disabled={isLoading}
                >
                    <FaCheckCircle />
                    {isLoading ? 'Procesando...' : 'Confirmar Pago y Programar Entrega'}
                </button>
            </div>
        </div>
    );
};

export default Payment;