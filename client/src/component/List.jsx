import React, { Component } from 'react';
import { useState } from 'react';
import {Button, ButtonToolbar} from 'react-bootstrap';

function SaveButton(todoes) {
    console.log(todoes)
    const arr = todoes.todoes.filter((todo) => todo.isDone)

    console.log("target:", arr.length)
    if(arr.length > 0) {
        return (
            <button size="lg" className="btn btn-outline-info btn-lg btn-block">Save</button>
        )        
    }
    return null;
}

export default function List() {
    const [allTodoes, setTodoes] = useState([
        {
            title: "",
            isDone: false,
        }
    ])

    const addTodo = () => {
        let prevTodesState = [...allTodoes];
        prevTodesState.push({
            title: "",
            isDone: false,
        });
        setTodoes(prevTodesState)
    };

    const deleteTodo = (evt) => {
        const id = evt.target.id
        const prevTodesState = [...allTodoes]
        prevTodesState.splice(id, 1)
        console.log(prevTodesState)
        setTodoes(prevTodesState)
    };

    const doCheck = (evt) =>  {
        const target = evt.target;
        const id = target.id;
        if(target.value === "") {
            return
        }
        let newTodoState = [...allTodoes]
        newTodoState[id].isDone = true
        setTodoes(newTodoState)
    }

    const removeCheck = (evt) => {
        const target = evt.target;
        const id = target.id;
        console.log(id)
        let newTodoState = [...allTodoes]
        newTodoState[id].isDone = false
        setTodoes(newTodoState)
    }

    const handleChange = (evt) => {
        const target = evt.target;
        const id = target.id;
        console.log(target)
        let newTodoState = [...allTodoes]
        newTodoState[id].title = target.value
        setTodoes(newTodoState)
    }

    console.log(allTodoes)
    return(
        <div className="row">
        <h2>TodoList</h2>
        <div className="d-grid gap-2">
            <Button variant="primary" size="lg" onClick={addTodo}>
            <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" fill="currentColor" className="plus-button bi bi-plus" viewBox="0 0 16 16" onClick={addTodo}>
                    <path d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4z"/>
            </svg>
            Add TODO
            </Button>
        </div>
          <div className="d-flex justify-content-start">
            <div className="flex-fill">
                {allTodoes.map((t, index) => {
                    if (!t.isDone) {
                        return(
                            <div key={index}>
                                <label className="list-group-item">
                                    <input className="form-check-input me-1" type="checkbox" checked={t.isDone} value={t.title} id={index} onChange={doCheck}/>
                                    <input className="input-text-box form-text" type="text" name="title" value={t.title} id={index} onChange={handleChange}/>
                                    <Button variant="danger" size="sm" onClick={deleteTodo} id={index} className="pull-right">Delete</Button>
                                </label>
                            </div>
                        )
                    }
                })}
            </div>
            <div className="list-group flex-fill">
                {allTodoes.map((t, index) => {
                    if (t.isDone) {
                        return(
                            <div key={index}>
                                <label className="list-group-item">
                                <input  className="form-check-input me-1"
                                        type="checkbox"
                                        checked={t.isDone}
                                        value={t.title}
                                        id={index}
                                        onChange={removeCheck}
                                />
                                {t.title}
                                </label>
                            </div>
                        )
                    }
                })}
            </div>
          </div>
        <SaveButton todoes={allTodoes}/>
        </div>
    )
}