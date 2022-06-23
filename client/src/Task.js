import TaskItem from './TaskItem';
import { useState, useEffect } from 'react';

// const tasks = [
//   {
//     id: 1,
//     text: 'Title',
//     day: 'DAY',
//     reminder: true,
//   },
//   {
//     id: 2,
//     text: 'Title2',
//     day: 'DAY2',
//     reminder: false,
//   },
// ];

const Tasks = () => {
  const [tasks, setTasks] = useState([]);

  const fetchTasks = async () => {
    const res = await fetch('http://localhost:5000/api/tasks');

    const tasks = await res.json();

    return tasks.data;
  };

  useEffect(() => {
    fetchTasks().then(tasks => {
      console.log(tasks);

      setTasks(tasks);
    });
  }, []);

  return (
    <ul>
      {tasks.map(task => (
        <TaskItem key={task.id} task={task} />
      ))}
    </ul>
  );
};

export default Tasks;
