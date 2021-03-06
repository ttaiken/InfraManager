﻿$(document).ready(function() {
  $("#table-servers").tabulator({
    pagination:"remote", 
    paginationSize: 6,
    ajaxURL: "http://192.168.6.1/infra/getservers",
    ajaxParams:{token:"ABC123"},
    tooltips:true,
    addRowPos:"top",
    history:true,
    movableColumns:true,
    //dataEdited:function(data){ alert(data[0].hostname)},
    
    cellEdited:function(cell){
        row = cell.getRow()
        alert(cell.getColumn().getField() + " - " + cell.getValue())
        alert(cell.getRow().getData().hostname)//+row.os +row.platform)
    },
    fitColumns:true,
    initialSort:[ {column:"hostname", dir:"asc"},],
    columns:[ 
        {title:"<input type='checkbox' name='checkall' />", field: "hostname",formatter:function(cell, formatterParams){
         return "<input type='checkbox' name='delinfo' value='" +
            cell.getValue()+ "' />";
        
        },width:50},
        {title:"hostname", field:"hostname", width:150,editor:"input",
          dataEdited:function(data){ alert(data[0].hostname)},rowEdit:function(id, data, row){alert(row)}},
        {title:"ip", field:"ip",editor:"input"},
        {title:"os", field:"os",editor:"input"},
        {title:"platform", field:"platform", align:"center",editor:"input"},
    ]
  });

//var data = $("#example-table").tabulator("getData");
//var row = $("#example-table").tabulator("getRow", 1); 
//$("#example-table").tabulator("clearData");
//$("#example-table").tabulator("updateOrAddRow", 3, {id:3, name:"steve", gender:"male"});
//$("#example-table").tabulator("deleteRow", 15);
$("#add-row").on("click", function(){
    $("#table-servers").tabulator("addRow",{} , true ); 
  });
$("#delete-row").on("click", function(){
  var items = [];
    item=$('[name="delinfo"]:checked').each(function(index,elem){
        items.push($(this).val());}); 
    window.alert(items);
    //$("#table-servers").tabulator("addRow",{} , true ); 
  });
$("#history-undo").on("click", function(){
    $("#table-servers").tabulator("undo");
});
$("#history-redo").on("click", function(){
    $("#table-servers").tabulator("redo");
});

});

