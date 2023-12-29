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
      <BTable striped bitemed="true" hover size="sm"  {...getTableProps()}>
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

export const ItemContent = (props) => {
    const [item, setItem] = React.useState(null);
    const [data, setData] = React.useState([]);
    const [selectDBSource, setSelectDBSource] = React.useState(null);
    const { register, handleSubmit} = useForm();
    
    const updateItemSelection = async (data) => {
      const selectedOrderId = data.orderId;; 
      let resJson = await getOrdersById(selectedOrderId);
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
        setItem(props.item);
    }, [props.item]);
    
    useEffect(() => {
        updateItemsWithTable();
        updateItemsOption();
    }, []);

    async function updateItemsOption() {
      const resJson = await getItems();
      resJson.forEach(element => {
          console.log(element);
      });
      console.log(resJson);
      let options = [];
      resJson.forEach(element => {
        options.push({id:element.order_id, value:element.order_id});
      });
      let itemsOption = {
        key:"orderId", 
        option:options, 
        label:"訂單ID"
      };
      console.log(itemsOption);
      setSelectDBSource(itemsOption);

  }

    function getItems(){
        const headers = new Headers()
        const options = {
            method: "GET",
            headers: headers,
            //mode: 'no-cors'
        };
        const uri = 'http://localhost:8080/items'
        
        return fetch(uri, options).then(response => response.json())
    }

    function getOrdersById(orderId){
      const headers = new Headers()
      const options = {
          method: "GET",
          headers: headers,
          //mode: 'no-cors'
      };
      const uri = 'http://localhost:8080/items/'+ orderId
      
      return fetch(uri, options).then(response => response.json())
  }

    async function updateItemsWithTable(){
        let resJson = await getItems();
        resJson.forEach(element => {
            console.log(element);
        });
        console.log(resJson);
        setData(resJson);
    }
    const columns = React.useMemo(
        () => [
          {
            Header: 'Item ID',
            columns: [
              {
                Header: 'Item ID',
                accessor: 'id',
              },
            ],
          },
          {
            Header: 'Order ID',
            columns: [
              {
                Header: 'Order ID',
                accessor:'order_id',
              },
            ],
          },
          {
            Header: 'Product ID',
            columns: [
              {
                Header: 'Product ID',
                accessor:'product_id',
              },
            ],
          },
          {
            Header: 'Is Shipped',
            columns: [
              {
                Header: 'Is Shipped',
                accessor:'is_shipped',
                Cell: ({ value }) => (value ? 'Yes' : 'No'), // Render 'Yes' or 'No' based on boolean value
              },
            ],
          }
        ],
        []
      )
    
    
    return (
      <div id="ItemContent-div">
        <p style={{ textAlign: 'center' }}>
          ItemContent
        </p>
        {item}
        <div>
        <Form onSubmit={handleSubmit(updateItemSelection)}>
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

