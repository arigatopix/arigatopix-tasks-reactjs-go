const TaskItem = ({ task }) => {
  return (
    <>
      <li>
        <p>{task.text}</p>
        <p>{task.day}</p>
      </li>
    </>
  );
};

export default TaskItem;
