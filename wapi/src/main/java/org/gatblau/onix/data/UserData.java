/*
    Onix Config Manager - Copyright (c) 2018-2020 by www.gatblau.org

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.

    Contributors to this project, hereby assign copyright in their code to the
    project, to be licensed under the same terms as the rest of the code.
*/

package org.gatblau.onix.data;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

import java.io.Serializable;

@ApiModel
public class UserData implements Serializable {
    private static final long serialVersionUID = 1L;

    private String key;
    private String name;
    private String email;
    private String pwd;
    private String salt;
    private String created;
    private String updated;
    private Integer version;
    private String changedBy;

    public UserData() {
    }

    public UserData(String key) {
        this.key = key;
    }

    public String getKey() {
        return key;
    }

    public void setKey(String key) {
        this.key = key;
    }

    @ApiModelProperty(
            position = 1,
            required = true,
            value = "The user name.")
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    @ApiModelProperty(
            position = 2,
            required = true,
            value = "The user email address.")
    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }
    
    @ApiModelProperty(
            position = 3,
            required = true,
            value = "The one-way encrypted password.")
    public String getPwd() {
        return pwd;
    }

    public void setPwd(String pwd) {
        this.pwd = pwd;
    }

    @ApiModelProperty(
            position = 4,
            required = false,
            value = "The date on which the item was created.",
            notes = "This value is always autogenerated by the service upon item creation and does not need to be specified. " +
                    "It is subsequently retrieved every time item information is requested.",
            allowEmptyValue = true
    )
    public String getCreated() {
        return created;
    }

    public void setCreated(String created) {
        this.created = created;
    }

    @ApiModelProperty(
            position = 5,
            required = false,
            value = "The date on which the item was last updated.",
            notes = "This value is always autogenerated by the service upon item updates and does not need to be specified. " +
                    "It is subsequently retrieved every time item information is requested."
    )
    public String getUpdated() {
        return updated;
    }

    public void setUpdated(String updated) {
        this.updated = updated;
    }

    @ApiModelProperty(
            position = 6,
            required = false,
            value = "The item version.",
            notes = "This value is always autogenerated by the service upon item creation or updates. " +
                    "It only needs to be specified if optimistic concurrency is required, for example by a user interface service."
    )
    public Integer getVersion() {
        return version;
    }

    public void setVersion(Integer version) {
        this.version = version;
    }

    public String getChangedBy() {
        return changedBy;
    }

    public void setChangedBy(String changedBy) {
        this.changedBy = changedBy;
    }

    public String getSalt() {
        return salt;
    }

    public void setSalt(String salt) {
        this.salt = salt;
    }
}