// Este componente de React muestra el historial de órdenes de gas realizadas por el usuario, permitiendo visualizar su estado y acceder a detalles.

import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { FaEye, FaBoxOpen } from 'react-icons/fa';
import '../Styles/dashboard.css';
import { useAuth } from '../../context/AuthContext';
import { orderService } from '../../services/ApiServices';


const OrdersPlaced = () => {
  const [orders, setOrders] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const { user } = useAuth();

  useEffect(() => {
    const fetchUserOrders = async () => {
      if (!user?.id) {
        setIsLoading(false);
        return;
      }
      try {
        const response = await orderService.getByUserId(user.id);
        setOrders(response.data || []);
      } catch (error) {
        console.error("Error al obtener las órdenes del usuario:", error);
        setOrders([]);
      } finally {
        setIsLoading(false);
      }
    };

    fetchUserOrders();
  }, [user]);

  const formatDate = (dateString) => {
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return new Date(dateString).toLocaleDateString('es-ES', options);
  };

  if (isLoading) {
    return <div className="loading-spinner">Cargando tus órdenes...</div>;
  }

  if (orders.length === 0) {
    return (
      <div className="empty-state">
        <FaBoxOpen size={50} style={{ marginBottom: '20px', color: 'var(--fruit-salad-400)' }} />
        <h2>Aún no has realizado ningún pedido</h2>
        <p>¡Empieza ahora y gestiona tus bombonas de gas fácilmente!</p>
        <Link to="/orders/new" className="btn-main-action">Crear mi primer pedido</Link>
      </div>
    );
  }

  return (
    <div className="dashboard-content">
      <header className="dashboard-header">
        <h1>Mis Órdenes</h1>
        <p>Aquí tienes un historial de todos los pedidos que has realizado.</p>
      </header>

      <div className="dashboard-card orders-table-card">
        <table className="orders-table">
          <thead>
            <tr>
              <th>ID del Pedido</th>
              <th>Fecha</th>
              <th>Estado</th>
              <th>Total</th>
              <th>Acciones</th>
            </tr>
          </thead>
          <tbody>
            {orders.map((order) => (
              <tr key={order.id}>
                <td><span className="order-id-highlight">ORD-{order.id}</span></td>
                <td>{formatDate(order.created_at)}</td>
                <td>
                  <span className={`status-badge status-${order.order_state.name.toLowerCase()}`}>
                    {order.order_state.name}
                  </span>
                </td>
                <td>${order.total_price.toFixed(2)}</td>
                <td>
                  <Link to={`/orders/${order.id}`} className="action-btn">
                    <FaEye /> Ver Detalles
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

export default OrdersPlaced;