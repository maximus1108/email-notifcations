import React, { useEffect, useState, useCallback } from "react";
import Product from './Product'
import Cart from './Cart'
import Checkout from './Checkout'

export default () => {

  const [ products, setProducts ] = useState([]);
  const [ selected, setSelected ] = useState(new Map());
  const [ checkoutActive, showCheckout ] = useState(false);

  const closeCheckout = useCallback(_ => showCheckout(false), []);
  const openCheckout = useCallback(_ => showCheckout(true), []);

  const addToCart = useCallback(product => {

    const existing = selected.get(product.id);

    if(!existing || existing.quantity < existing.stock) {
    
      const quantity = 1 + (existing ? existing.quantity : 0);

      const newSelected = new Map(selected).set(product.id, {
        ...product,
        quantity
      });

      setSelected(newSelected);

    }

  }, [selected]);

  const sumbitOrder = useCallback(user => {

    const products = Array.from(selected.values());
    const body = JSON.stringify({
      ...user,
      products
    });

    fetch('http://localhost:8088/', {
      method: 'POST',
      body
    })
    .then(res => res.json())
    .then(json => console.log(json))
    .catch(err => console.log('err post', err))
  }, [selected]);

  useEffect(_ => {
    fetch('http://localhost:8088/')
      .then(res => res.json())
      .then(products => setProducts(products));
  }, []);

  return (
    <>
      <h1>Products</h1>
      <div className='product-list'>
        {
          products.map((product) => (
            <Product
              product={ product }
              handleClick={ addToCart }
              key={ product.id }
            />
          ))
        }
      </div>
      <Cart
        products={ Array.from(selected.values()) }
        onClick={ openCheckout }
      />
      {
        checkoutActive &&
          <Checkout
            onCancel={ closeCheckout }
            onSubmit={ sumbitOrder }
          />
      }
    </>
  );
}
