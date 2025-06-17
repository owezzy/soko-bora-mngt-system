import {createClientToken} from "./connect.module";
import {CustomersService} from "../proto/customerspb/api_pb";

export const CustomerGrpcService = createClientToken(CustomersService);

// Additional client tokens representing Connect services could be added here
