import React, {useEffect} from "react";
import BTable from 'react-bootstrap/Table';
import { useTable } from 'react-table'
import { useForm } from "react-hook-form";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Stack from 'react-bootstrap/Stack';
import { InputGroup } from "react-bootstrap";



function Table({ columns, data }) {
    // Use the state and functions returned from useTable to build your UI
    const {
      getTableProps,
      getTableBodyProps,
      headerGroups,
      rows,
      prepareRow,
    } = useTable({
      columns,
      data,
    })
  
    // Render the UI for your table
    return (
      <BTable striped bordered hover size="sm"  {...getTableProps()}>
        <thead>
          {headerGroups.map(headerGroup => (
            <tr {...headerGroup.getHeaderGroupProps()}>
              {headerGroup.headers.map(column => (
                <th {...column.getHeaderProps()}>{column.render('Header')}</th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody {...getTableBodyProps()}>
          {rows.map((row, i) => {
            prepareRow(row)
            return (
              <tr {...row.getRowProps()}>
                {row.cells.map(cell => {
                  return <td {...cell.getCellProps()}>{cell.render('Cell')}</td>
                })}
              </tr>
            )
          })}
        </tbody>
      </BTable>
    )
  }

export const OrderContent = (props) => {
    const [order, setOrder] = React.useState(null);
    const [data, setData] = React.useState([]);
    const [selectDBSource, setSelectDBSource] = React.useState(null);
    const { register, handleSubmit} = useForm();
    
    const updateOrderSelection = async (data) => {
      const selectedCustomerId = data.customerId;; 
      let resJson = await getOrdersById(selectedCustomerId);
      console.log(resJson);
      setData([resJson]);
      if(resJson.length == 0){
          alert("查無資料");
      }
  }

    const generateFormComps = (obj) => {
      if (!obj) return null;
      const { key, option, label } = obj;

      return (
          <InputGroup key={key}>
              <InputGroup.Text id={key}>{label}</InputGroup.Text>
              <Form.Select className="w-25" {...register(key)}>
                  {option.map(({ id, value }) => <option value={value} key={value}>{id}</option>)}
              </Form.Select>
          </InputGroup>
      );

  };

    useEffect(() => {
        setOrder(props.order);
    }, [props.order]);
    
    useEffect(() => {
        updateOrdersWithTable();
        updateOrdersOption();
    }, []);

    async function updateOrdersOption() {
      const resJson = await getOrders();
      resJson.forEach(element => {
          console.log(element);
      });
      console.log(resJson);
      let options = [];
      resJson.forEach(element => {
        options.push({id:element.customer_id, value:element.customer_id});
      });
      let ordersOption = {
        key:"customerId", 
        option:options, 
        label:"客戶ID"
      };
      console.log(ordersOption);
      setSelectDBSource(ordersOption);

  }

    function getOrders(){
        const headers = new Headers()
        const options = {
            method: "GET",
            headers: headers,
            //mode: 'no-cors'
        };
        const uri = 'http://localhost:8080/orders'
        
        return fetch(uri, options).then(response => response.json())
    }

    function getOrdersById(customerId){
      const headers = new Headers()
      const options = {
          method: "GET",
          headers: headers,
          //mode: 'no-cors'
      };
      const uri = 'http://localhost:8080/orders/'+ customerId
      
      return fetch(uri, options).then(response => response.json())
  }

    async function updateOrdersWithTable(){
        let resJson = await getOrders();
        resJson.forEach(element => {
            console.log(element);
        });
        console.log(resJson);
        setData(resJson);
    }

    const columns = React.useMemo(
        () => [
          {
            Header: 'Order ID',
            columns: [
              {
                Header: 'Order ID',
                accessor: 'id',
              },
            ],
          },
          {
            Header: 'Customer ID',
            columns: [
              {
                Header: 'Customer ID',
                accessor:'customer_id',
              },
            ],
          },
          {
            Header: 'Is Paid',
            columns: [
              {
                Header: 'Is Paid',
                accessor:'is_paid',
                Cell: ({ value }) => (value ? 'Yes' : 'No'),
              },
            ],
          }
        ],
        []
      )
    
    
    return (
      <div id="OrderContent-div">
        <p style={{ textAlign: 'center' }}>
          OrderContent
        </p>
        {order}
        <div>
        <Form onSubmit={handleSubmit(updateOrderSelection)}>
                <Stack direction="horizontal" gap={3}>
                    {generateFormComps(selectDBSource)}
                    <Button type="submit" className="btn btn-dark">查詢</Button>
                </Stack>
            </Form>
        </div>
        <div>
          <Table columns={columns} data={data} />
        </div>
      </div>
    
    );
}
    
