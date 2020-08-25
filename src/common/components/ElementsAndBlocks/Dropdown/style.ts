import {css} from "emotion";
import colors from "../../../styles/colors";

export const Dropdown = css`
    position: absolute;
    right: 0px;
    background-color: ${colors.white};
    z-index: 2;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 4px 10px 0px;
    border: 1px solid ${colors.greenActive};
    border-radius: 4px;
    box-sizing: border-box;
`;
