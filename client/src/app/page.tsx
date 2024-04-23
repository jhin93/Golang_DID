'use client'

import styled from 'styled-components';
import { useState } from 'react';

export default function Page() {
  const [userId, setUserId] = useState('');

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
      event.preventDefault();
      console.log('User ID Submitted:', userId);
  };

  return (
      <Container>
          <h1>Welcome to DID System</h1>
          <Form onSubmit={handleSubmit}>
              <Input
                  type="text"
                  value={userId}
                  onChange={(e) => setUserId(e.target.value)}
                  placeholder="Enter your user ID"
                  required
              />
              <Button type="submit">Submit</Button>
          </Form>
      </Container>
  );
}

const Container = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    background-color: #f0f2f5;
    color: black;
`;

const Form = styled.form`
    display: flex;
    flex-direction: column;
    width: 400px; /* 너비 증가 */
    // background: white;
    padding: 30px; /* 패딩 증가 */
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
`;

const Input = styled.input`
    margin-bottom: 15px; /* 여백 증가 */
    padding: 15px; /* 패딩 증가 */
    border: 1px solid #ccc;
    border-radius: 4px;
    background: white;
    color: black;
`;

const Button = styled.button`
    padding: 15px 30px; /* 버튼 패딩 증가 */
    // color: white;
    background-color: #0070f3;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s;

    &:hover {
        background-color: #0056b3;
    }
`;
