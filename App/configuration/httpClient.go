package configuration

import 


func ProccessRequest( id,name,verb,url,body,headers,err string){


}


func ResolveRequest(url,body,headers,verb string) (
	switch verb {
		case model.httpVerb.GET:
			 doGet(url,headers)
		case model.httpVerb.POST:
			 doGet(url,headers)
		case model.httpVerb.PUT:
			 doGet(url,headers)
		case model.httpVerb.DELETE:
			 doGet(url,headers)	
		}
}


func doGet(url,headers){

}
