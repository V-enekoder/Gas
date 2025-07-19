// Este componente de React muestra el historial de entregas de gas realizadas al usuario, recuperando los datos desde la API y presentando un estado vacío si no hay entregas.

import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { FaEye, FaShippingFast } from 'react-icons/fa';
import '../Styles/dashboard.css';
import { useAuth } from '../../context/AuthContext';
import { deliveryService } from '../../services/ApiServices';


const Deliveries = () => {
  const [deliveries, setDeliveries] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const { user } = useAuth();

  useEffect(() => {
    const fetchUserDeliveries = async () => {
      if (!user?.id) {
        setIsLoading(false);
        return;
      }
      setIsLoading(true);

      try {
        const response = await deliveryService.getByUserId(user.id);

        if (Array.isArray(response.data)) {
          setDeliveries(response.data);
        } else if (response.data && typeof response.data === 'object' && response.data.id) {
          setDeliveries([response.data]);
        } else {
          setDeliveries([]);
        }

      } catch (error) {
        console.error("Error al obtener las entregas (Backend Error):", error);
        setDeliveries([]);
      } finally {
        setIsLoading(false);
      }
    };

    fetchUserDeliveries();
  }, [user]);


  if (isLoading) {
    return <div className="loading-spinner">Cargando tus entregas...</div>;
  }

  if (deliveries.length === 0) {
    return (
      <div className="empty-state">
        <FaShippingFast size={50} style={{ marginBottom: '20px', color: 'var(--fruit-salad-400)' }} />
        <h2>No tienes entregas registradas</h2>
        <p>Cuando uno de tus pedidos sea despachado, aparecerá aquí.</p>
        <Link to="/orders" className="btn-main-action">Ver mis pedidos</Link>
      </div>
    );
  }

  return (
    <div className="dashboard-content">
      <header className="dashboard-header">
        <h1>Mis Entregas</h1>
        <p>Consulta el historial de todas las entregas que han sido despachadas.</p>
      </header>

      <div className="dashboard-card orders-table-card">
        <table className="orders-table">
          <thead>
            <tr>
              <th>ID de Entrega</th>
              <th>ID de Orden</th>
              <th>Total</th>
              <th>Acciones</th>
            </tr>
          </thead>
          <tbody>
            {deliveries.map((delivery) => (
              <tr key={delivery.id}>
                <td><span className="order-id-highlight">DEL-{delivery.id}</span></td>
                <td>ORD-{delivery.order.id}</td>
                <td>${delivery.total_price.toFixed(2)}</td>
                <td>
                  <Link to={`/deliveries/${delivery.id}`} className="action-btn">
                    <FaEye /> Ver Recibo
                  </Link>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Deliveries;