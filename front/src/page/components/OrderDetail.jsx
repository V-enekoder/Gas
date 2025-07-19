// Este archivo define la página de 'Detalles del Pedido', que busca y muestra la información de un pedido específico a partir de su ID en la URL.

import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { orderService } from '../../services/ApiServices';
import { FaDollarSign, FaCalendarAlt, FaTruckLoading, FaBoxOpen } from 'react-icons/fa';
import '../Styles/components.css';

const OrderDetail = () => {
  const { orderId } = useParams(); 
  const navigate = useNavigate();
  
  const [order, setOrder] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchOrderDetails = async () => {
      if (!orderId) {
        setError("ID de orden no proporcionado.");
        setIsLoading(false);
        return;
      }
      try {
        const response = await orderService.getById(orderId);
        setOrder(response.data);
      } catch (err) {
        console.error("Error al obtener los detalles de la orden:", err);
        setError("No se pudieron cargar los detalles de la orden.");
      } finally {
        setIsLoading(false);
      }
    };
    fetchOrderDetails();
  }, [orderId]);

  const formatDate = (dateString) => {
    if (!dateString) return 'N/A';
    const options = { year: 'numeric', month: 'long', day: 'numeric', hour: '2-digit', minute: '2-digit' };
    return new Date(dateString).toLocaleDateString('es-ES', options);
  };

  const handleProceedToPayment = () => {
    navigate(`/orders/${orderId}/payment`);
  };

  if (isLoading) {
    return <div className="loading-spinner">Cargando detalles de la orden...</div>;
  }

  if (!order) {
    return (
      <div className="dashboard-content">
        <header className="dashboard-header"><h1>Error</h1></header>
        <p className="error-message">{error || "No se encontró la orden."}</p>
      </div>
    );
  }

  const isPendingPayment = order.order_state.name.toLowerCase() === 'pendiente de pago';

  return (
    <div className="dashboard-content">
      <header className="dashboard-header">
        <h1>Detalles del Pedido ORD-{order.id}</h1>
        <p>Toda la información sobre tu pedido.</p>
      </header>

      <div className="dashboard-card order-detail-card">
        <div className="detail-section">
          <h2>Información General</h2>
          <div className="detail-grid">
            <div className="detail-item">
              <FaDollarSign />
              <span>Total Pagado:</span>
              <strong>${order.total_price.toFixed(2)}</strong>
            </div>
            <div className="detail-item">
              <FaCalendarAlt />
              <span>Fecha de Creación:</span>
              <strong>{formatDate(order.created_at)}</strong>
            </div>
            <div className="detail-item">
              <FaTruckLoading />
              <span>Estado del Pedido:</span>
              <span className={`status-badge status-${order.order_state.name.toLowerCase().replace(' ', '-')}`}>
                {order.order_state.name}
              </span>
            </div>
          </div>
        </div>

        <div className="detail-section">
          <h2>Productos del Pedido</h2>
          <table className="order-details-table">
            <thead>
              <tr>
                <th>Tipo de Bombona</th>
                <th>Cantidad</th>
                <th>Precio Unitario</th>
                <th>Subtotal</th>
              </tr>
            </thead>
            <tbody>
              {order.order_details.map(detail => (
                <tr key={detail.id}>
                  <td>{detail.type_cylinder.name}</td>
                  <td>{detail.quantity}</td>
                  <td>${detail.price.toFixed(2)}</td>
                  <td>${(detail.price * detail.quantity).toFixed(2)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
        
        {isPendingPayment && (
          <div className="payment-action-section">
            <p>Este pedido está pendiente de pago.</p>
            <button onClick={handleProceedToPayment} className="btn-main-action">
              <FaDollarSign /> Proceder al Pago Ahora
            </button>
          </div>
        )}

        {error && <p className="error-message">{error}</p>}
      </div>
    </div>
  );
};

export default OrderDetail;