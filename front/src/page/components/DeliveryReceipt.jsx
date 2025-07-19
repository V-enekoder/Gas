// Este componente de React muestra un recibo de entrega detallado obtenido desde una API.

import React, { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { deliveryService } from '../../services/ApiServices';
import '../Styles/components.css';

const DeliveryReceipt = () => {
  const { id } = useParams(); 
  const [delivery, setDelivery] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchDelivery = async () => {
      try {
        const response = await deliveryService.getById(id);
        setDelivery(response.data);
      } catch (error) {
        console.error("Error al obtener el recibo:", error);
        setDelivery(null);
      } finally {
        setIsLoading(false);
      }
    };

    fetchDelivery();
  }, [id]);

  const formatDate = (dateStr) => {
    if (!dateStr) return 'N/A';
    return new Date(dateStr).toLocaleDateString('es-ES', { day: 'numeric', month: 'long', year: 'numeric' });
  };

  if (isLoading) return <div className="loading-spinner">Cargando recibo...</div>;
  if (!delivery) return <div className="error-message">No se pudo cargar el recibo.</div>;

  return (
    <div className="receipt-container">
      <div className="receipt-card">
        <h1>Recibo de Entrega</h1>
        <p><strong>ID de Entrega:</strong> DEL-{delivery.id}</p>
        <p><strong>ID de Orden:</strong> ORD-{delivery.order.id}</p>

        <table className="receipt-table">
          <thead>
            <tr>
              <th>Tipo de Cilindro</th>
              <th>Cantidad</th>
              <th>Subtotal</th>
            </tr>
          </thead>
          <tbody>
            {delivery.delivery_details.map((detail) => (
              <tr key={detail.id}>
                <td>{detail.type_cylinder.name}</td>
                <td>{detail.quantity}</td>
                <td>${(detail.quantity * (delivery.total_price / totalCylinders(delivery.delivery_details))).toFixed(2)}</td>
              </tr>
            ))}
          </tbody>
          <tfoot>
            <tr>
              <td colSpan="2"><strong>Total</strong></td>
              <td><strong>${delivery.total_price.toFixed(2)}</strong></td>
            </tr>
          </tfoot>
        </table>

        <div className="receipt-footer">
          <Link to="/deliveries" className="btn-main-action">Volver a mis entregas</Link>
        </div>
      </div>
    </div>
  );
};

const totalCylinders = (details) => {
  return details.reduce((sum, d) => sum + d.quantity, 0);
};

export default DeliveryReceipt;
