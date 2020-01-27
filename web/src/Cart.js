import React from 'react';
import CartItem from './CartItem';

export default ({ products, onClick }) => (
  <div className='cart'>
    <h2>Cart</h2>
    <div>
      <h3>Item</h3>
      <h3>Qty</h3>
    </div>
    {
      products.map(({ id, name, quantity }) =>
        <CartItem
          key={ id }
          name={ name }
          quantity={ quantity }
        />
      )
    }
    {
      !!products.length &&
        <button
          className='cart-btn'
          onClick={onClick}
        >
          Checkout
        </button>
    }
  </div>
);